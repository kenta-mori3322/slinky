package config_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/skip-mev/slinky/oracle/config"
)

func TestAPIConfig(t *testing.T) {
	testCases := []struct {
		name        string
		config      config.APIConfig
		expectedErr bool
	}{
		{
			name: "good config with api enabled",
			config: config.APIConfig{
				Enabled:          true,
				Timeout:          time.Second,
				Interval:         time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				Name:             "test",
				URL:              "http://test.com",
			},
			expectedErr: false,
		},
		{
			name: "good config with api disabled",
			config: config.APIConfig{
				Enabled: false,
			},
			expectedErr: false,
		},
		{
			name: "bad config with no timeout",
			config: config.APIConfig{
				Enabled:          true,
				Interval:         time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				Name:             "test",
				URL:              "http://test.com",
			},
			expectedErr: true,
		},
		{
			name: "bad config with no interval",
			config: config.APIConfig{
				Enabled:          true,
				Timeout:          time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				Name:             "test",
				URL:              "http://test.com",
			},
			expectedErr: true,
		},
		{
			name: "bad config with no max queries",
			config: config.APIConfig{
				Enabled:  true,
				Timeout:  time.Second,
				Interval: time.Second,
				Name:     "test",
				URL:      "http://test.com",
			},
			expectedErr: true,
		},
		{
			name: "bad config with no name",
			config: config.APIConfig{
				Enabled:          true,
				Timeout:          time.Second,
				Interval:         time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				URL:              "http://test.com",
			},
			expectedErr: true,
		},
		{
			name: "bad config with no url / endpoint",
			config: config.APIConfig{
				Enabled:          true,
				Timeout:          time.Second,
				Interval:         time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				Name:             "test",
			},
			expectedErr: true,
		},
		{
			name: "bad config with no reconnect timeout",
			config: config.APIConfig{
				Enabled:    true,
				Timeout:    time.Second,
				Interval:   time.Second,
				MaxQueries: 1,
				Name:       "test",
				URL:        "http://test.com",
			},
			expectedErr: true,
		},
		{
			name: "good config with endpoints instead of url",
			config: config.APIConfig{
				Enabled:          true,
				Timeout:          time.Second,
				Interval:         time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				Name:             "test",
				Endpoints:        []config.Endpoint{{URL: "http://test.com"}},
			},
			expectedErr: false,
		},
		{
			name: "bad config with atomic + batch-size",
			config: config.APIConfig{
				Enabled:          true,
				Timeout:          time.Second,
				Interval:         time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				Name:             "test",
				Endpoints:        []config.Endpoint{{URL: "http://test.com"}},
				Atomic:           true,
				BatchSize:        1,
			},
			expectedErr: true,
		},
		{
			name: "good config with batchSize",
			config: config.APIConfig{
				Enabled:          true,
				Timeout:          time.Second,
				Interval:         time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				Name:             "test",
				Endpoints:        []config.Endpoint{{URL: "http://test.com"}},
				BatchSize:        1,
			},
			expectedErr: false,
		},
		{
			name: "bad config with invalid endpoint",
			config: config.APIConfig{
				Enabled:          true,
				Timeout:          time.Second,
				Interval:         time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				Name:             "test",
				Endpoints:        []config.Endpoint{
					{
						URL: "http://test.com",
						Authentication: config.Authentication{
							Enabled: true,
						},
					},
				},
				BatchSize:        1,	
			},
			expectedErr: true,
		},
		{
			name: "good config with valid endpoint",
			config: config.APIConfig{
				Enabled:          true,
				Timeout:          time.Second,
				Interval:         time.Second,
				ReconnectTimeout: time.Second,
				MaxQueries:       1,
				Name:             "test",
				Endpoints:        []config.Endpoint{
					{
						URL: "http://test.com",
						Authentication: config.Authentication{
							Enabled: true,
							HTTPHeaderAPIKey: "test",
						},
					},
				},
				BatchSize:        1,	
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.ValidateBasic()
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
