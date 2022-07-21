# 更改日志 | CHANGE LOG

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
