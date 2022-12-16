package http

import "testing"

func TestHealthCheck(t *testing.T) {
	cases := map[string]struct {
		appHealthy      bool
		dbHealthy       bool
		expectedHealthy bool
	}{
		"All Okay": {appHealthy: true, dbHealthy: true, expectedHealthy: true},
		"App down": {appHealthy: false, dbHealthy: true, expectedHealthy: false},
		"Db Down":  {appHealthy: true, dbHealthy: false, expectedHealthy: false},
		"All down": {appHealthy: false, dbHealthy: false, expectedHealthy: false},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			h := Health{App: test.appHealthy, Db: test.dbHealthy}
			actualHealth := h.isHealthy()
			if actualHealth != test.expectedHealthy {
				t.Errorf("health is %t, expected heathy to be %t given app health is %v",
					actualHealth,
					test.expectedHealthy,
					h)
			}
		})
	}
}
