package data

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// static_app_js reads file data from disk. It returns an error on failure.
func static_app_js() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/app.js"
	name := "static/app.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_graphspace_jpg reads file data from disk. It returns an error on failure.
func static_graphspace_jpg() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/graphspace.jpg"
	name := "static/graphspace.jpg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_index_html reads file data from disk. It returns an error on failure.
func static_index_html() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/index.html"
	name := "static/index.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_jquery_1_11_2_js reads file data from disk. It returns an error on failure.
func static_jquery_1_11_2_js() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/jquery-1.11.2.js"
	name := "static/jquery-1.11.2.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_jquery_1_11_2_min_js reads file data from disk. It returns an error on failure.
func static_jquery_1_11_2_min_js() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/jquery-1.11.2.min.js"
	name := "static/jquery-1.11.2.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_jquery_linedtextarea_css reads file data from disk. It returns an error on failure.
func static_jquery_linedtextarea_css() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/jquery-linedtextarea.css"
	name := "static/jquery-linedtextarea.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_jquery_linedtextarea_js reads file data from disk. It returns an error on failure.
func static_jquery_linedtextarea_js() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/jquery-linedtextarea.js"
	name := "static/jquery-linedtextarea.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_jquery_ui_css reads file data from disk. It returns an error on failure.
func static_jquery_ui_css() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/jquery-ui.css"
	name := "static/jquery-ui.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_jquery_ui_js reads file data from disk. It returns an error on failure.
func static_jquery_ui_js() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/jquery-ui.js"
	name := "static/jquery-ui.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_diagonals_thick_18_b81900_40x40_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_diagonals_thick_18_b81900_40x40_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_diagonals-thick_18_b81900_40x40.png"
	name := "static/images/ui-bg_diagonals-thick_18_b81900_40x40.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_diagonals_thick_20_666666_40x40_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_diagonals_thick_20_666666_40x40_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_diagonals-thick_20_666666_40x40.png"
	name := "static/images/ui-bg_diagonals-thick_20_666666_40x40.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_flat_0_aaaaaa_40x100_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_flat_0_aaaaaa_40x100_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_flat_0_aaaaaa_40x100.png"
	name := "static/images/ui-bg_flat_0_aaaaaa_40x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_flat_10_000000_40x100_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_flat_10_000000_40x100_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_flat_10_000000_40x100.png"
	name := "static/images/ui-bg_flat_10_000000_40x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_flat_75_ffffff_40x100_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_flat_75_ffffff_40x100_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_flat_75_ffffff_40x100.png"
	name := "static/images/ui-bg_flat_75_ffffff_40x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_glass_100_f6f6f6_1x400_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_glass_100_f6f6f6_1x400_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_glass_100_f6f6f6_1x400.png"
	name := "static/images/ui-bg_glass_100_f6f6f6_1x400.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_glass_100_fdf5ce_1x400_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_glass_100_fdf5ce_1x400_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_glass_100_fdf5ce_1x400.png"
	name := "static/images/ui-bg_glass_100_fdf5ce_1x400.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_glass_55_fbf9ee_1x400_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_glass_55_fbf9ee_1x400_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_glass_55_fbf9ee_1x400.png"
	name := "static/images/ui-bg_glass_55_fbf9ee_1x400.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_glass_65_ffffff_1x400_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_glass_65_ffffff_1x400_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_glass_65_ffffff_1x400.png"
	name := "static/images/ui-bg_glass_65_ffffff_1x400.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_glass_75_dadada_1x400_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_glass_75_dadada_1x400_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_glass_75_dadada_1x400.png"
	name := "static/images/ui-bg_glass_75_dadada_1x400.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_glass_75_e6e6e6_1x400_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_glass_75_e6e6e6_1x400_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_glass_75_e6e6e6_1x400.png"
	name := "static/images/ui-bg_glass_75_e6e6e6_1x400.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_glass_95_fef1ec_1x400_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_glass_95_fef1ec_1x400_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_glass_95_fef1ec_1x400.png"
	name := "static/images/ui-bg_glass_95_fef1ec_1x400.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_gloss_wave_35_f6a828_500x100_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_gloss_wave_35_f6a828_500x100_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_gloss-wave_35_f6a828_500x100.png"
	name := "static/images/ui-bg_gloss-wave_35_f6a828_500x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_highlight_soft_100_eeeeee_1x100_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_highlight_soft_100_eeeeee_1x100_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_highlight-soft_100_eeeeee_1x100.png"
	name := "static/images/ui-bg_highlight-soft_100_eeeeee_1x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_highlight_soft_75_cccccc_1x100_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_highlight_soft_75_cccccc_1x100_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_highlight-soft_75_cccccc_1x100.png"
	name := "static/images/ui-bg_highlight-soft_75_cccccc_1x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_bg_highlight_soft_75_ffe45c_1x100_png reads file data from disk. It returns an error on failure.
func static_images_ui_bg_highlight_soft_75_ffe45c_1x100_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-bg_highlight-soft_75_ffe45c_1x100.png"
	name := "static/images/ui-bg_highlight-soft_75_ffe45c_1x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_icons_222222_256x240_png reads file data from disk. It returns an error on failure.
func static_images_ui_icons_222222_256x240_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-icons_222222_256x240.png"
	name := "static/images/ui-icons_222222_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_icons_228ef1_256x240_png reads file data from disk. It returns an error on failure.
func static_images_ui_icons_228ef1_256x240_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-icons_228ef1_256x240.png"
	name := "static/images/ui-icons_228ef1_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_icons_2e83ff_256x240_png reads file data from disk. It returns an error on failure.
func static_images_ui_icons_2e83ff_256x240_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-icons_2e83ff_256x240.png"
	name := "static/images/ui-icons_2e83ff_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_icons_454545_256x240_png reads file data from disk. It returns an error on failure.
func static_images_ui_icons_454545_256x240_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-icons_454545_256x240.png"
	name := "static/images/ui-icons_454545_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_icons_888888_256x240_png reads file data from disk. It returns an error on failure.
func static_images_ui_icons_888888_256x240_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-icons_888888_256x240.png"
	name := "static/images/ui-icons_888888_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_icons_cd0a0a_256x240_png reads file data from disk. It returns an error on failure.
func static_images_ui_icons_cd0a0a_256x240_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-icons_cd0a0a_256x240.png"
	name := "static/images/ui-icons_cd0a0a_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_icons_ef8c08_256x240_png reads file data from disk. It returns an error on failure.
func static_images_ui_icons_ef8c08_256x240_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-icons_ef8c08_256x240.png"
	name := "static/images/ui-icons_ef8c08_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_icons_ffd27a_256x240_png reads file data from disk. It returns an error on failure.
func static_images_ui_icons_ffd27a_256x240_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-icons_ffd27a_256x240.png"
	name := "static/images/ui-icons_ffd27a_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_ui_icons_ffffff_256x240_png reads file data from disk. It returns an error on failure.
func static_images_ui_icons_ffffff_256x240_png() (*asset, error) {
	path := "/home/sig/go/graphspace/src/github.com/sigmonsays/graphspace/static/images/ui-icons_ffffff_256x240.png"
	name := "static/images/ui-icons_ffffff_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"static/app.js": static_app_js,
	"static/graphspace.jpg": static_graphspace_jpg,
	"static/index.html": static_index_html,
	"static/jquery-1.11.2.js": static_jquery_1_11_2_js,
	"static/jquery-1.11.2.min.js": static_jquery_1_11_2_min_js,
	"static/jquery-linedtextarea.css": static_jquery_linedtextarea_css,
	"static/jquery-linedtextarea.js": static_jquery_linedtextarea_js,
	"static/jquery-ui.css": static_jquery_ui_css,
	"static/jquery-ui.js": static_jquery_ui_js,
	"static/images/ui-bg_diagonals-thick_18_b81900_40x40.png": static_images_ui_bg_diagonals_thick_18_b81900_40x40_png,
	"static/images/ui-bg_diagonals-thick_20_666666_40x40.png": static_images_ui_bg_diagonals_thick_20_666666_40x40_png,
	"static/images/ui-bg_flat_0_aaaaaa_40x100.png": static_images_ui_bg_flat_0_aaaaaa_40x100_png,
	"static/images/ui-bg_flat_10_000000_40x100.png": static_images_ui_bg_flat_10_000000_40x100_png,
	"static/images/ui-bg_flat_75_ffffff_40x100.png": static_images_ui_bg_flat_75_ffffff_40x100_png,
	"static/images/ui-bg_glass_100_f6f6f6_1x400.png": static_images_ui_bg_glass_100_f6f6f6_1x400_png,
	"static/images/ui-bg_glass_100_fdf5ce_1x400.png": static_images_ui_bg_glass_100_fdf5ce_1x400_png,
	"static/images/ui-bg_glass_55_fbf9ee_1x400.png": static_images_ui_bg_glass_55_fbf9ee_1x400_png,
	"static/images/ui-bg_glass_65_ffffff_1x400.png": static_images_ui_bg_glass_65_ffffff_1x400_png,
	"static/images/ui-bg_glass_75_dadada_1x400.png": static_images_ui_bg_glass_75_dadada_1x400_png,
	"static/images/ui-bg_glass_75_e6e6e6_1x400.png": static_images_ui_bg_glass_75_e6e6e6_1x400_png,
	"static/images/ui-bg_glass_95_fef1ec_1x400.png": static_images_ui_bg_glass_95_fef1ec_1x400_png,
	"static/images/ui-bg_gloss-wave_35_f6a828_500x100.png": static_images_ui_bg_gloss_wave_35_f6a828_500x100_png,
	"static/images/ui-bg_highlight-soft_100_eeeeee_1x100.png": static_images_ui_bg_highlight_soft_100_eeeeee_1x100_png,
	"static/images/ui-bg_highlight-soft_75_cccccc_1x100.png": static_images_ui_bg_highlight_soft_75_cccccc_1x100_png,
	"static/images/ui-bg_highlight-soft_75_ffe45c_1x100.png": static_images_ui_bg_highlight_soft_75_ffe45c_1x100_png,
	"static/images/ui-icons_222222_256x240.png": static_images_ui_icons_222222_256x240_png,
	"static/images/ui-icons_228ef1_256x240.png": static_images_ui_icons_228ef1_256x240_png,
	"static/images/ui-icons_2e83ff_256x240.png": static_images_ui_icons_2e83ff_256x240_png,
	"static/images/ui-icons_454545_256x240.png": static_images_ui_icons_454545_256x240_png,
	"static/images/ui-icons_888888_256x240.png": static_images_ui_icons_888888_256x240_png,
	"static/images/ui-icons_cd0a0a_256x240.png": static_images_ui_icons_cd0a0a_256x240_png,
	"static/images/ui-icons_ef8c08_256x240.png": static_images_ui_icons_ef8c08_256x240_png,
	"static/images/ui-icons_ffd27a_256x240.png": static_images_ui_icons_ffd27a_256x240_png,
	"static/images/ui-icons_ffffff_256x240.png": static_images_ui_icons_ffffff_256x240_png,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"app.js": &_bintree_t{static_app_js, map[string]*_bintree_t{
		}},
		"graphspace.jpg": &_bintree_t{static_graphspace_jpg, map[string]*_bintree_t{
		}},
		"images": &_bintree_t{nil, map[string]*_bintree_t{
			"ui-bg_diagonals-thick_18_b81900_40x40.png": &_bintree_t{static_images_ui_bg_diagonals_thick_18_b81900_40x40_png, map[string]*_bintree_t{
			}},
			"ui-bg_diagonals-thick_20_666666_40x40.png": &_bintree_t{static_images_ui_bg_diagonals_thick_20_666666_40x40_png, map[string]*_bintree_t{
			}},
			"ui-bg_flat_0_aaaaaa_40x100.png": &_bintree_t{static_images_ui_bg_flat_0_aaaaaa_40x100_png, map[string]*_bintree_t{
			}},
			"ui-bg_flat_10_000000_40x100.png": &_bintree_t{static_images_ui_bg_flat_10_000000_40x100_png, map[string]*_bintree_t{
			}},
			"ui-bg_flat_75_ffffff_40x100.png": &_bintree_t{static_images_ui_bg_flat_75_ffffff_40x100_png, map[string]*_bintree_t{
			}},
			"ui-bg_glass_100_f6f6f6_1x400.png": &_bintree_t{static_images_ui_bg_glass_100_f6f6f6_1x400_png, map[string]*_bintree_t{
			}},
			"ui-bg_glass_100_fdf5ce_1x400.png": &_bintree_t{static_images_ui_bg_glass_100_fdf5ce_1x400_png, map[string]*_bintree_t{
			}},
			"ui-bg_glass_55_fbf9ee_1x400.png": &_bintree_t{static_images_ui_bg_glass_55_fbf9ee_1x400_png, map[string]*_bintree_t{
			}},
			"ui-bg_glass_65_ffffff_1x400.png": &_bintree_t{static_images_ui_bg_glass_65_ffffff_1x400_png, map[string]*_bintree_t{
			}},
			"ui-bg_glass_75_dadada_1x400.png": &_bintree_t{static_images_ui_bg_glass_75_dadada_1x400_png, map[string]*_bintree_t{
			}},
			"ui-bg_glass_75_e6e6e6_1x400.png": &_bintree_t{static_images_ui_bg_glass_75_e6e6e6_1x400_png, map[string]*_bintree_t{
			}},
			"ui-bg_glass_95_fef1ec_1x400.png": &_bintree_t{static_images_ui_bg_glass_95_fef1ec_1x400_png, map[string]*_bintree_t{
			}},
			"ui-bg_gloss-wave_35_f6a828_500x100.png": &_bintree_t{static_images_ui_bg_gloss_wave_35_f6a828_500x100_png, map[string]*_bintree_t{
			}},
			"ui-bg_highlight-soft_100_eeeeee_1x100.png": &_bintree_t{static_images_ui_bg_highlight_soft_100_eeeeee_1x100_png, map[string]*_bintree_t{
			}},
			"ui-bg_highlight-soft_75_cccccc_1x100.png": &_bintree_t{static_images_ui_bg_highlight_soft_75_cccccc_1x100_png, map[string]*_bintree_t{
			}},
			"ui-bg_highlight-soft_75_ffe45c_1x100.png": &_bintree_t{static_images_ui_bg_highlight_soft_75_ffe45c_1x100_png, map[string]*_bintree_t{
			}},
			"ui-icons_222222_256x240.png": &_bintree_t{static_images_ui_icons_222222_256x240_png, map[string]*_bintree_t{
			}},
			"ui-icons_228ef1_256x240.png": &_bintree_t{static_images_ui_icons_228ef1_256x240_png, map[string]*_bintree_t{
			}},
			"ui-icons_2e83ff_256x240.png": &_bintree_t{static_images_ui_icons_2e83ff_256x240_png, map[string]*_bintree_t{
			}},
			"ui-icons_454545_256x240.png": &_bintree_t{static_images_ui_icons_454545_256x240_png, map[string]*_bintree_t{
			}},
			"ui-icons_888888_256x240.png": &_bintree_t{static_images_ui_icons_888888_256x240_png, map[string]*_bintree_t{
			}},
			"ui-icons_cd0a0a_256x240.png": &_bintree_t{static_images_ui_icons_cd0a0a_256x240_png, map[string]*_bintree_t{
			}},
			"ui-icons_ef8c08_256x240.png": &_bintree_t{static_images_ui_icons_ef8c08_256x240_png, map[string]*_bintree_t{
			}},
			"ui-icons_ffd27a_256x240.png": &_bintree_t{static_images_ui_icons_ffd27a_256x240_png, map[string]*_bintree_t{
			}},
			"ui-icons_ffffff_256x240.png": &_bintree_t{static_images_ui_icons_ffffff_256x240_png, map[string]*_bintree_t{
			}},
		}},
		"index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
		}},
		"jquery-1.11.2.js": &_bintree_t{static_jquery_1_11_2_js, map[string]*_bintree_t{
		}},
		"jquery-1.11.2.min.js": &_bintree_t{static_jquery_1_11_2_min_js, map[string]*_bintree_t{
		}},
		"jquery-linedtextarea.css": &_bintree_t{static_jquery_linedtextarea_css, map[string]*_bintree_t{
		}},
		"jquery-linedtextarea.js": &_bintree_t{static_jquery_linedtextarea_js, map[string]*_bintree_t{
		}},
		"jquery-ui.css": &_bintree_t{static_jquery_ui_css, map[string]*_bintree_t{
		}},
		"jquery-ui.js": &_bintree_t{static_jquery_ui_js, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

