# 更改日志 | CHANGE LOG

- 提高块跳转超链接与子窗口的 `z-index` 层级 | Raise the 'z-index' level of block jump hyperlinks and popover windows.
- 将子窗口预览功能默认触发时间调整为 `1000ms` | Adjust the time of the default trigger time of the popover preview window to `1000ms`.
- 使用 `getDoc` API 查询文档的 DOM | Use the `getDoc` API to query the DOM of the document.
- 为文档添加面包屑 | Add a breadcrumb for the document.
- 鼠标悬浮块超链接时高亮对应的块 | Highlight the block when the mouse is hovering over the block hyperlink.
- 修复跨域超链接预览功能 | Fix the cross-domain hyperlink preview function.

## v0.1.1 / 2022-07-28

- [v0.1.0 ... v0.1.1](https:///github.com/Zuoqiu-Yingyi/siyuan-publish/compare/v0.1.0...v1.1.1)
- 为每个块添加跳转超链接 | Add a link to each block to jump to it.
- 文档内超链接支持鼠标悬浮小窗预览 | Hyperlinks within the document support popover preview.
- 修复悬浮预览窗口拖拽失效问题 | Fixed the issue that dragging and dropping feature on the popover preview window failed.
- 抽象并封装拖动功能 | Abstract and encapsulate the dragfeature.
- 悬浮预览窗口支持使用鼠标拖动调整尺寸 | Popover preview window supports using the mouse to drag to adjust the size.
- 支持自定义预览超链接目标的小窗口样式 | Supports custom small window styles for previewing hyperlinks target.
- 支持设置默认文档访问权限 | Supports setting the default document access permission.

## v0.1.0 / 2022-07-25

- 初始化项目 | Initialize project.
- 添加主题基础配色文件路径配置选项 | Add theme base color file path configuration option.
- 支持使用 URL 参数 `theme` 指定主题模式 | Support use URL parameter `theme` to specify theme mode.
  - `theme=light`: 明亮主题模式 | Light theme mode.
  - `theme=dark`: 暗黑主题模式 | Dark theme mode.
  - `theme=auto`: 根据浏览器环境自动选择主题模式 | Auto select theme mode by browser environment.
- 支持使用文档块属性 `custom-publish-access` 设置访问权限 | Support use document block attribute `custom-publish-access` to set access permission.
  - `public`: 公开访问 | Public access.
  - `private`: 私有访问 | Private access.
  - `protected`: 受保护访问 | Protected access.
- 支持在配置文件 `*.config.toml` 中设置代码块渲染样式 | Support set code block render style in `*.config.toml` file.
- 支持在配置文件 `*.config.toml` 中设置字号 | Support set font size in `*.config.toml` file.
- 支持块引用超链接跳转时携带参数 | Support jump to block reference with parameters.
- 使用结构体整理异常状态响应 | Use structure to organize exception status response.
- 支持自定义首页路径列表与首页重定向的 URL | Support custom homepage path list and homepage redirect URL.
- 添加资源文件动态加载模式 | Add resource file dynamic loading mode.
- 解析文档内嵌入的链接 | Parse embedded links in document.
- 添加资源文件动态缓存加载模式 | Add resource file dynamic cache loading mode.
- 资源文件动态加载模式响应转发切换为 `DataFromReader` 方法 | Resource file dynamic loading mode response forwarding switch to `DataFromReader` method.
- 设计数据模型 | Design data model.
- 输出日志文件 | Output log file.
- 使用 MVC 框架重构项目结构 | Use MVC framework to rebuild project structure.
- 实现文档页面加载的动态缓存模式 | Implement document page loading dynamic cache mode.
- 实现文档页面加载的静态加载模式 | Implement document page loading static mode.
- 添加启动时重置静态资源目录配置选项 | Add reset static resource directory configuration option when booting.
- 添加文档标题 | Add document title.
- 移动目录 `/app/` 至 `/kernel/app/` | Move directory `/app/` to `/kernel/app/`.
- 添加文档图标, 文档标签与文档题头图 | Add document icon, document tags and document title image.
- 添加 CI/CD 工具配置 | Add CI/CD tool configuration.
- 文档页面静态加载模式初始化时同时加载受保护的文档 | Document page static loading mode initialization time loads protected document at the same time.
- 内核启动时创建日志目录 | Kernel booting time create log directory.
- 使用 `github.com/glebarez/sqlite` 作为 GROM 的 SQLite 驱动 | Use `github.com/glebarez/sqlite` as GROM SQLite driver.
- 添加字体列表配置选项 | Add font list configuration option.
- 添加 `favicon.ico` 文件 | Add `favicon.ico` file.
- 添加中文文档 | Add Chinese document.
- 解析文档内嵌入的思源链接(`siyuan://`) | Parse embedded siyuan links (`siyuan://`) in document.
