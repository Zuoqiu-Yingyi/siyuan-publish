# 中文文档

## 快速开始

### 要求

* 设备已安装思源笔记或可以访问思源笔记伺服服务
* 在 [Releases](https://github.com/Zuoqiu-Yingyi/siyuan-publish/releases) 中下载与设备 `操作系统` 及 `架构` 一致的最新的应用发行版
* 如无特殊说明, 下文中 `内核` 特指本应用的可执行程序, `思源内核` 指提供思源笔记伺服功能的应用程序

### 安装

* 将与设备 `操作系统` 及 `架构` 一致的最新应用安装包 `*.zip` 移动到安装目录并解压即可

### 项目结构

```plaintext
.
│  CHANGELOG.md // 更改日志
│  default.config.toml // 默认配置文件
│  LICENSE // 开源许可
│  publish.exe // 内核
│  README.md // 自述文件
│
├─app // web 应用静态文件
│  ├─src // web 资源文件
│  │  │   favicon.ico // 站点图标
│  │  │
│  │  ├─css // 样式文件
│  │  │      font.css // 默认字体大小样式
│  │  │      publish.css // 自定义样式
│  │  │
│  │  └─js // 脚本文件
│  │          config.js // 自定义配置
│  │          publish.js // 自定义渲染
│  │
│  └─templates // 模板
│          background.html // 题头模板
│          block.html // 文档模板
│          config.html // 配置模板
│          error.html // 错误模板
│          render.html // 渲染器模板
│          title.html // 文档标题模板
│
├─docs // 介绍
│       README-en.md // 英文介绍
│       README-zh-SC.md // 中文介绍
│
└─temp // 运行时的临时文件目录(运行时自动生成)
   │    publish.db // 数据库文件
   │
   ├─logs // 运行日志目录
   └─static // 静态资源文件临时目录
      ├─appearance // 外观资源目录(对应思源笔记工作区 /conf/appearance/ 目录)
      ├─aeests // 文档资源文件目录(对应思源笔记工作区 /data/assets/ 目录)
      ├─emojis // 表情资源文件目录(对应思源笔记工作区 /data/emojis/ 目录)
      ├─export // 导出资源文件目录(对应思源笔记工作区 /temp/export/ 目录)
      ├─stage // 渲染工具资源文件目录(对应思源笔记安装 /resource/stage/ 目录)
      └─widgets // 挂件资源文件目录(对应思源笔记工作区 /data/widgets/ 目录)
```

### 设置文档访问权限

1. 打开思源笔记
2. 找到想要发布的文档
3. 为文档添加自定义块属性, 属性名为 `publish-access`, 属性值可为 `public`, `protected` 或 `private` 三者中的一个

   * <kbd>publish-access</kbd>:`public`

     * 该文档及其任意深度的下级文档都可被公开访问, 除非其下级文档设置了自定义块属性<kbd>publish-access</kbd>:`protected` 或 `private`
   * <kbd>publish-access</kbd>:`protected`**(TODO)**

     * 该文档及其任意深度的下级文档在通过鉴权后都可访问, 除非其下级文档设置了自定义块属性<kbd>publish-access</kbd>:`public` 或 `private`
   * <kbd>publish-access</kbd>:`private`

     * 该文档及其任意深度的下级文档都不可访问, 除非其下级文档设置了自定义块属性 <kbd>publish-access</kbd>:`public` 或 `protected`
   * 若一篇文档及其上级文档都没有设置自定义块<kbd>publish-access</kbd>, 那么该文档不可访问

### 自定义配置文件

* 默认配置文件为应用根目录下的 `default.config.toml` 文件

  * 该文件为内核默认加载的配置文件, 若想手动指定配置文件路径, 需要在启动内核时传入参数 `--config <自定义配置文件的路径>`

    * 例如想手动指定自定义配置文件为 `安装目录/custom.config.toml`, 那么使用命令 `./publish --config ./custom.config.toml` 带参数启动内核
* 配置文件格式为 TOML, 详情请参考 [TOML：为人而生的配置文件格式](https://toml.io/cn/)

  * 完整规范: [TOML v1.0.0](https://toml.io/cn/latest)
* 初次运行内核前需要首先设置如下字段

  1. `Server.Port`: Web 服务访问端口, 默认为 `80` 端口
  2. `Server.Mode.Page`: 文档页面加载模式, 字段值可以设置为如下三者之一

     * `dynamic`: 动态加载模式, 设置为该模式时内核将实时向思源内核请求数据

       * 内核实时判断访问的文档是否是公开的
       * 内核不使用数据库缓存访问控制列表(Access Control List, ACL), 因此通过编辑文档自定义属性<kbd>publish-access</kbd>更改文档的访问控制权限时可以实时生效
       * 不能关闭思源内核, 思源内核关闭或思源内核服务无法访问时发布内容无法正常访问
     * `cache`: 动态缓存模式, 设置为该模式时内核将先从数据库中查询缓存内容, 若没有查询到缓存内容才会向思源内核请求数据并写入缓存

       * 内核会在启动时向思源内核请求数据建立 ACL, 并将 ACL 写入数据库中, 因此通过编辑文档自定义属性<kbd>publish-access</kbd>更改文档的访问控制权限时不能实时生效, 必须在重启内核时才能更新 ACL

         * 注: 如果设置了 `Database.Reset = false`, 那么重新启动内核时也不会重建 ACL
       * 文档内容与文档关联的访问控制权限在第一次访问后也会写入数据库中, 因此当更改已缓存的文档内容或移动已缓存的文档后, 发布的内容也不会实时更新, 也同样需要重启内核才能清除缓存内容

         * 注: 如果设置了 `Database.Reset = false`, 那么重新启动内核时也不会清除缓存
       * 不能关闭思源内核, 思源内核关闭或思源内核服务无法访问时, 缓存的内容可以正常访问, 缓存外的内容无法正常访问
     * `static`: 静态加载模式, 设置为该模式时内核会在启动时将所有发布的文档信息写入数据库中

       * 内核会在启动时向思源内核请求数据建立 ACL, 并将 ACL 写入数据库中, 因此通过编辑文档自定义属性<kbd>publish-access</kbd>更改文档的访问控制权限时不能实时生效, 必须在重启内核时才能更新 ACL

         * 注: 如果设置了 `Database.Reset = false`, 那么重新启动内核时也不会重建 ACL
       * 文档内容与文档关联的访问控制权限在内核启动时也会同时写入数据库中, 因此内核启动后更改已发布文档内容或移动已发布文档时, 发布的内容也不会实时更新, 可通过重启内核更新数据库

         * 注: 如果设置了 `Database.Reset = false`, 那么重新启动内核时也不会更新数据库
       * 思源内核关闭或无法访问时, 发布的文档也可以正常访问
  3. `Server.Mode.File`: 资源文件加载模式, 其字段值可以设置为如下三者之一

     * `dynamic`: 动态加载模式, 该模式下向内核请求资源文件时内核实时将请求转发至思源内核

       * 由于思源 `assets` 目录下的资源在发布时只要知道资源名称即可公开访问, 因此该选项具有一定的安全风险
       * 思源内核关闭或不能访问时资源无法加载
     * `cache`: 动态缓存模式, 该模式下向内核请求资源文件时内核先判断对应静态资源目录(缓存目录)是否存在该文件, 若不存在则向思源内核请求下载该文件并保存到对应的静态资源目录(缓存目录)

       * 由于思源 `assets` 目录下的资源在发布时只要知道资源名称即可公开访问, 因此该选项具有一定的安全风险
       * 思源内核关闭或不能访问时已缓存的资源可以正常加载, 未缓存的资源不能加载
     * `static`: 静态加载模式, 该模式下向内核将静态资源目录映射为静态文件服务

       * 该模式下只能访问静态资源目录下存在的资源文件, 因此较安全
       * 思源内核关闭或不能访问时, 静态资源目录中的文件都能正常加载
  4. `Server.Index`: 自定义站点首页路径, 设置好后访问该路径列表中的某一项时都会重定向至设置的首页 URL
  5. `Siyuan.Server`: 思源内核服务地址, 用于加载想要发布的资源
  6. `Siyuan.Server.Token`: 思源内核服务的访问令牌, 若思源内核服务启用了访问授权码则需要配置该选项, 可以在思源<kbd>设置</kbd>><kbd>关于</kbd>><kbd>API Token</kbd>中复制
  7. `Render.File.Style.Base`: 思源基础样式文件的 URL, 建议设置为思源安装目录 `/resource/stage/build/mobile/` 下的 CSS 文件访问 URL, 例如该 CSS 文件名为 `base.2ad8890755ebaf8cc6d3.css`, 那么该选项设置为 `/stage/build/mobile/base.2ad8890755ebaf8cc6d3.css`
* 保存配置文件

### 启动内核

* 内核为安装目录根目录中文件名为 `publish` 的可执行程序, 运行时如果使用与内核文件处于同一目录, 文件名为 `default.config.toml` 的配置文件, 可以不带参数启动, 否则需要使用参数 `--config` 指定资源文件路径

### 访问发布的内容

* 访问首页

  * `http(s)://host:port`
* 访问 ID 为 `20200812220555-lj3enxa` 的文档

  * `http(s)://host:port/block/20200812220555-lj3enxa`
* 访问 ID 为 `20210428212840-859h45j` 的块所在的文档

  * `http(s)://host:port/block/20210428212840-859h45j`
* 访问 ID 为 `20210428212840-859h45j` 的块所在的文档并定位到该块

  * `http(s)://host:port/block?id=20210428212840-859h45j`
  * `http(s)://host:port/block/20200812220555-lj3enxa?id=20210428212840-859h45j`
* 使用亮色主题访问 ID 为 `20200812220555-lj3enxa` 的文档

  * `http(s)://host:port/block/20200812220555-lj3enxa?theme=light`
* 使用暗色主题访问 ID 为 `20200812220555-lj3enxa` 的文档

  * `http(s)://host:port/block/20200812220555-lj3enxa?theme=dark`
