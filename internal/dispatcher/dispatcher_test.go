package dispatcher_test

import (
	"testing"
	"web-crawler/internal/dispatcher"
	"web-crawler/internal/model"
)

type mockHandler struct {
	name        string
	shouldMatch bool
	returnBlock model.Block
	returnError error
}

func (m *mockHandler) Match(html string) bool {
	return m.shouldMatch
}

func (m *mockHandler) Extract(html, pageURL string) (model.Block, error) {
	return m.returnBlock, m.returnError
}

func (m *mockHandler) Type() string {
	return m.name
}

func TestDispatch(t *testing.T) {
	originalHandlers := model.GetAllHandlers()

	tests := []struct {
		name        string
		mockMap     map[string]model.BlockHandler
		expected    []model.Block
		expectError bool
	}{
		{
			name: "Multiple handlers match and extract",
			mockMap: map[string]model.BlockHandler{
				"mock1": &mockHandler{
					name:        "mock1",
					shouldMatch: true,
					returnBlock: model.Block{
						Type:     "mock1",
						HTML:     "<div>mock1</div>",
						PageURL:  "http://site",
						Found:    "true",
						Accuracy: "1.0",
					},
				},
				"mock2": &mockHandler{
					name:        "mock2",
					shouldMatch: true,
					returnBlock: model.Block{
						Type:     "mock2",
						HTML:     "<div>mock2</div>",
						PageURL:  "http://site",
						Found:    "true",
						Accuracy: "1.0",
					},
				},
			},
			expected: []model.Block{
				{
					Type:     "mock1",
					HTML:     "<div>mock1</div>",
					PageURL:  "http://site",
					Found:    "true",
					Accuracy: "1.0",
				},
				{
					Type:     "mock2",
					HTML:     "<div>mock2</div>",
					PageURL:  "http://site",
					Found:    "true",
					Accuracy: "1.0",
				},
			},
			expectError: false,
		},
		{
			name: "No handler matches",
			mockMap: map[string]model.BlockHandler{
				"non-matching": &mockHandler{
					name:        "non-matching",
					shouldMatch: false,
				},
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace handlers
			overrideHandlerMap(tt.mockMap)
			defer overrideHandlerMap(originalHandlers)

			blocks, err := dispatcher.Dispatch("<html>mock</html>", "http://site")
			if (err != nil) != tt.expectError {
				t.Fatalf("Unexpected error state: %v", err)
			}

			if !equalBlocks(blocks, tt.expected) {
				t.Errorf("Unexpected result:\ngot:  %+v\nwant: %+v", blocks, tt.expected)
			}
		})
	}
}

func equalBlocks(a, b []model.Block) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// overrideHandlerMap заменяет глобальный map handlers через reflection (hacky, но безопасно внутри теста)
func overrideHandlerMap(newMap map[string]model.BlockHandler) {
	// Unsafe но в рамках теста корректно: пересоздаем map в пакете model
	for k := range model.GetAllHandlers() {
		delete(model.GetAllHandlers(), k)
	}
	for _, v := range newMap {
		model.RegisterHandler(v)
	}
}
