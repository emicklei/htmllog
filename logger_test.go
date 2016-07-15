package htmllog

import "testing"

func TestBasic(t *testing.T) {
	h, _ := New("basic.html")
	for i := 0; i < 10; i++ {
		h.Infof("here i <am>")
		h.Warnf("didn't i warn you")
		h.Debugf("about to throw error")
		h.Errorf("catch the error")
	}
}

func TestLimited(t *testing.T) {
	h, _ := New("limit.html")
	h.Configure("", DefaultStyle, 4, LogEvent)
	defer func() {
		h.Configure("", DefaultStyle, 200, LogEvent)
	}()
	h.Infof("123456789")
}

func TestReset(t *testing.T) {
	h, _ := New("reset.html")
	h.Infof("before reset")
	if err := h.Reset(); err != nil {
		t.Fatal(err)
	}
	h.Infof("after reset")
}

func TestScrolling(t *testing.T) {
	h, _ := New("scrolling.html")
	for i := 0; i < 100; i++ {
		h.Infof("here i <am>")
		h.Warnf("didn't i warn you")
		h.Debugf("about to throw error")
		h.Errorf("catch the error")
	}
}
