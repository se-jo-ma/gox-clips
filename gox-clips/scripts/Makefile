# vim: tabstop=8
CLIPS_VERSION		?= 6.42
CLIPS_SOURCE_URL	?= "https://sourceforge.net/projects/clipsrules/files/CLIPS/6.4.2/clips_core_source_642.zip"
LIBS_DIR		?= $(PWD)/libs
DIST_DIR		?= $(PWD)/dist
MAKEFILE_NAME		?= makefile
SHARED_INCLUDE_DIR	?= /usr/local/include
SHARED_LIBRARY_DIR	?= /usr/local/lib
TARGET_ARCH		?= $(shell uname -m)
LINUX_LDLIBS		?= -lm -lrt

.PHONY: clips install clean

all: clips_source clips

clips_source:
	curl --output clips.zip --location --url $(CLIPS_SOURCE_URL)
	unzip -jo clips.zip -d clips_source

clips: clips_source
	mkdir -p $(LIBS_DIR)
	$(MAKE) -f $(MAKEFILE_NAME) -C clips_source				\
		CFLAGS="-std=c99 -O3 -fno-strict-aliasing -fPIC"		\
		LDLIBS="$(LINUX_LDLIBS)"
	ld -G clips_source/*.o -o $(LIBS_DIR)/libclips.so

build: clips
	mkdir -p $(DIST_DIR)

repair: export LD_LIBRARY_PATH := $LD_LIBRARY_PATH:$(LIBS_DIR)
repair: build

install: clips build repair
	install -d $(SHARED_INCLUDE_DIR)/
	install -m 644 clips_source/clips.h $(SHARED_INCLUDE_DIR)/
	install -d $(SHARED_INCLUDE_DIR)/clips
	install -m 644 clips_source/*.h $(SHARED_INCLUDE_DIR)/clips/
	install -d $(SHARED_LIBRARY_DIR)/
	install -m 644 clips_source/libclips.so					\
	 	$(SHARED_LIBRARY_DIR)/libclips.so.$(CLIPS_VERSION)
	ln -sf $(SHARED_LIBRARY_DIR)/libclips.so.$(CLIPS_VERSION)		\
	 	$(SHARED_LIBRARY_DIR)/libclips.so.6
	ln -sf $(SHARED_LIBRARY_DIR)/libclips.so.$(CLIPS_VERSION)		\
	 	$(SHARED_LIBRARY_DIR)/libclips.so
	-ldconfig -n -v $(SHARED_LIBRARY_DIR)

clean:
	-rm clips.zip
	-rm -fr clips_source build dist clipspy.egg-info