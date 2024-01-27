package deep

import (
	"reflect"
	"time"
)

var defaultDeep = deep{ignoreTypes: map[reflect.Type]struct{}{
	reflect.TypeOf(time.Time{}): {},
}}

type deep struct {
	ignoreTypes map[reflect.Type]struct{}
}

func (d *deep) Copy(dst reflect.Value, src reflect.Value) {
	switch src.Kind() {
	case reflect.Interface:
		if !src.IsNil() {
			srcElem := src.Elem()
			dstElem := reflect.New(srcElem.Type()).Elem()
			d.Copy(dstElem, srcElem)
			dst.Set(dstElem)
		}
	case reflect.Map:
		if !src.IsNil() {
			dst.Set(reflect.MakeMap(src.Type()))
			for _, key := range src.MapKeys() {
				srcKeyElem := key
				dstKeyElem := reflect.New(srcKeyElem.Type()).Elem()
				d.Copy(dstKeyElem, srcKeyElem)
				srcItemElem := src.MapIndex(key)
				dstItemElem := reflect.New(srcItemElem.Type()).Elem()
				d.Copy(dstItemElem, srcItemElem)
				dst.SetMapIndex(dstKeyElem, dstItemElem)
			}
		}
	case reflect.Pointer:
		if !src.IsNil() {
			dst.Set(reflect.New(src.Elem().Type()))
			d.Copy(dst.Elem(), src.Elem())
		}
	case reflect.Slice:
		if !src.IsNil() {
			length := src.Len()
			dst.Set(reflect.MakeSlice(src.Type(), length, src.Cap()))
			for idx := 0; idx < length; idx++ {
				d.Copy(dst.Index(idx), src.Index(idx))
			}
		}
	case reflect.Struct:
		if src.IsZero() {
			return
		}
		srcTyp := src.Type()
		if _, ok := d.ignoreTypes[srcTyp]; ok {
			return
		}
		dst.Set(src)
		num := src.NumField()
		for idx := 0; idx < num; idx++ {
			if srcTyp.Field(idx).IsExported() {
				d.Copy(dst.Field(idx), src.Field(idx))
			}
		}
	default:
		dst.Set(src)
	}
}

func CopyIgnore[T any](src T, ignores ...any) (dst T) {
	ignoreTypes := map[reflect.Type]struct{}{}
	for _, ignore := range ignores {
		ignoreTypes[reflect.TypeOf(ignore)] = struct{}{}
	}
	srcRv := reflect.ValueOf(&src)
	dstRv := reflect.ValueOf(&dst)
	custom := deep{ignoreTypes: ignoreTypes}
	custom.Copy(dstRv.Elem(), srcRv.Elem())
	return
}

func Copy[T any](src T) (dst T) {
	srcRv := reflect.ValueOf(&src)
	dstRv := reflect.ValueOf(&dst)
	defaultDeep.Copy(dstRv.Elem(), srcRv.Elem())
	return
}
