OBJS := _deps/cimgui/cimgui.o
OBJS += _deps/cimgui/imgui/imgui.o
OBJS += _deps/cimgui/imgui/imgui_demo.o
OBJS += _deps/cimgui/imgui/imgui_draw.o
OBJS += _deps/cimgui/imgui/imgui_tables.o
OBJS += _deps/cimgui/imgui/imgui_widgets.o
OBJS += _deps/cimgui/imgui/backends/imgui_impl_glfw.o
OBJS += _deps/cimgui/imgui/backends/imgui_impl_opengl3.o

HOST_TARGET := libcimgui_$(shell go env GOOS)_$(shell go env GOARCH).a

CXXFLAGS := -Wall -O2 -fno-exceptions -fno-rtti
CXXFLAGS += -I _deps/cimgui/imgui -DIMGUI_IMPL_API='extern "C" '

AR := ar -rc

%.o : %.cpp
	$(CXX) -c $(CXXFLAGS) $< -o $@

.PHONY: all
all: $(HOST_TARGET) cimgui.h cimgui_impl.h imgui.h

$(HOST_TARGET): $(OBJS)
	$(AR) $(HOST_TARGET) $(OBJS)

cimgui.h: _deps/cimgui/cimgui.h
	cp -f _deps/cimgui/cimgui.h cimgui.h

cimgui_impl.h: _deps/cimgui/generator/output/cimgui_impl.h
	cp -f _deps/cimgui/generator/output/cimgui_impl.h cimgui_impl.h

imgui.h: _deps/cimgui/imgui/imgui.h
	cp -f _deps/cimgui/imgui/imgui.h imgui.h

.PHONY: clean
clean:
	rm -f $(HOST_TARGET) $(OBJS) cimgui.h cimgui_impl.h imgui.h
