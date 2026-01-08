package browser

import (
	"github.com/puter/web-browse-agent/internal/tool"
)

// RegisterBrowserTools registers all browser tools with the given registry
func RegisterBrowserTools(registry *tool.Registry, ctx *Context) {
	// Navigation tools
	registry.Register(NewNavigateTool(ctx))
	registry.Register(NewNavigateBackTool(ctx))

	// Interaction tools
	registry.Register(NewClickTool(ctx))
	registry.Register(NewFillTool(ctx))
	registry.Register(NewTypeTool(ctx))
	registry.Register(NewPressKeyTool(ctx))
	registry.Register(NewHoverTool(ctx))
	registry.Register(NewSelectOptionTool(ctx))

	// Page inspection tools
	registry.Register(NewSnapshotTool(ctx))
	registry.Register(NewScreenshotTool(ctx))
	registry.Register(NewEvaluateTool(ctx))

	// Tab management tools
	registry.Register(NewTabsListTool(ctx))
	registry.Register(NewTabNewTool(ctx))
	registry.Register(NewTabSelectTool(ctx))
	registry.Register(NewTabCloseTool(ctx))

	// Wait tools
	registry.Register(NewWaitTool(ctx))

	// Dialog and file tools
	registry.Register(NewHandleDialogTool(ctx))
	registry.Register(NewFileUploadTool(ctx))
}

// AllBrowserTools returns all browser tools for a given context
// This is useful when you need to manage tools separately from a registry
func AllBrowserTools(ctx *Context) []tool.Tool {
	return []tool.Tool{
		NewNavigateTool(ctx),
		NewNavigateBackTool(ctx),
		NewClickTool(ctx),
		NewFillTool(ctx),
		NewTypeTool(ctx),
		NewPressKeyTool(ctx),
		NewHoverTool(ctx),
		NewSelectOptionTool(ctx),
		NewSnapshotTool(ctx),
		NewScreenshotTool(ctx),
		NewEvaluateTool(ctx),
		NewTabsListTool(ctx),
		NewTabNewTool(ctx),
		NewTabSelectTool(ctx),
		NewTabCloseTool(ctx),
		NewWaitTool(ctx),
		NewHandleDialogTool(ctx),
		NewFileUploadTool(ctx),
	}
}

// BrowserToolNames returns the names of all browser tools
func BrowserToolNames() []string {
	return []string{
		"browser_navigate",
		"browser_navigate_back",
		"browser_click",
		"browser_fill",
		"browser_type",
		"browser_press_key",
		"browser_hover",
		"browser_select_option",
		"browser_snapshot",
		"browser_screenshot",
		"browser_evaluate",
		"browser_tabs_list",
		"browser_tab_new",
		"browser_tab_select",
		"browser_tab_close",
		"browser_wait",
		"browser_handle_dialog",
		"browser_file_upload",
	}
}
