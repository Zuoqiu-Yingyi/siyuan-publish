# English Documentation

> This document was translated by ChatGPT-4

## Quick Start

### Requirements

* The device has installed SiYuan Note or can access SiYuan Note server service
* Download the latest app distribution for the device `Operating System` and `Architecture` in [Releases](https://github.com/Zuoqiu-Yingyi/siyuan-publish/releases)
* Unless otherwise stated, `kernel` below refers to the executable program of this application, and `SiYuan kernel` refers to the application program providing SiYuan Note server functionality

### Installation

* Move the latest app installation package `*.zip` for the device `Operating System` and `Architecture` to the installation directory and unzip it

### Project Structure

```
.
│  CHANGELOG.md // Change log
│  default.config.toml // Default configuration file
│  LICENSE // Open source license
│  publish.exe // Kernel
│  README.md // README file
│
├─app // Web app static files
│  ├─src // Web resource files
│  │  │   favicon.ico // Site icon
│  │  │
│  │  ├─css // Style files
│  │  │      font.css // Default font size style
│  │  │      publish.css // Custom styles
│  │  │
│  │  └─js // Script files
│  │     │   after.js // Post-render run script
│  │     │   before.js // Pre-render run script
│  │     │   context.js // Plugin context
│  │     │   plugin.js // Plugin initialization
│  │     │   render.js // Document rendering
│  │     │   utils.js // Utility components
│  │     │
│  │     └─plugins // Render plugins
│  │
│  └─templates // Templates
│          background.html // Header template
│          block.html // Document template
│          config.html // Configuration template
│          error.html // Error template
│          render.html // Renderer template
│          title.html // Document title template
│
├─docs // Introduction
│       README-en.md // English introduction
│       README-zh-SC.md // Chinese introduction
│
├─scripts // Scripts
│       dev-build.ps1 // Developer build script
│       dev-run.ps1 // Developer run script
│
└─temp // Temporary file directory at runtime (auto-generated at runtime)
   │    publish.db // Database file
   │
   ├─logs // Running log directory
   └─static // Static resource file temporary directory
      ├─appearance // Appearance resource directory (corresponding to SiYuan Note workspace /conf/appearance/ directory)
      ├─aeests // Document resource file directory (corresponding to SiYuan Note workspace /data/assets/ directory)
      ├─emojis // Emotion resource file directory (corresponding to SiYuan Note workspace /data/emojis/ directory)
      ├─export // Export resource file directory (corresponding to SiYuan Note workspace /temp/export/ directory)
      ├─stage // Rendering tool resource file directory (corresponding to SiYuan Note installation /resource/stage/ directory)
      └─widgets // Widget resource file directory (corresponding to SiYuan Note workspace /data/widgets/ directory)
```

* The configuration file format is TOML, for more information, please refer to [TOML: A human-friendly configuration file format](https://toml.io/cn/)

  * Full specification: [TOML v1.0.0](https://toml.io/cn/latest)
* Before running the kernel for the first time, you need to set the following fields:

  1. `Server.Port`: Web service access port, the default is port `80`
  2. `Server.Mode.Page`: Document page loading mode, field value can be set to one of the following three:

      * `dynamic`: Dynamic loading mode, when set to this mode, the kernel will request data from the SiYuan kernel in real-time

        * The kernel determines in real-time whether the accessed document is public
        * The kernel does not use the database cache Access Control List (ACL), so changing the access control permission of a document by editing the custom attribute<kbd>publish-access</kbd> can take effect in real-time
        * SiYuan kernel cannot be turned off, the published content cannot be accessed normally when the SiYuan kernel is turned off or the SiYuan kernel service is inaccessible
      * `cache`: Dynamic cache mode, when set to this mode, the kernel will first query the cache content from the database, and if the cache content is not found, it will request data from the SiYuan kernel and write it to the cache

        * The kernel will request data from the SiYuan kernel to establish an ACL at startup and write the ACL to the database, so when editing the custom attribute<kbd>publish-access</kbd> to change the access control permission of a document, it cannot take effect in real-time and must be restarted in the kernel to update the ACL

          * Note: If `Database.Reset = false` is set, restarting the kernel will not rebuild the ACL
        * The document content and the access control permission associated with the document will be written to the database after the first access, so after changing the cached document content or moving the cached document, the published content will not be updated in real-time, and the kernel needs to be restarted to clear the cache content

          * Note: If `Database.Reset = false` is set, restarting the kernel will not clear the cache
        * SiYuan kernel cannot be turned off, the cached content can be accessed normally when the SiYuan kernel is turned off or the SiYuan kernel service is inaccessible, and the content outside the cache cannot be accessed normally
      * `static`: Static loading mode, when set to this mode, the kernel will write all published document information to the database at startup

        * The kernel will request data from the SiYuan kernel to establish an ACL at startup and write the ACL to the database, so when editing the custom attribute<kbd>publish-access</kbd> to change the access control permission of a document, it cannot take effect in real-time and must be restarted in the kernel to update the ACL

          * Note: If `Database.Reset = false` is set, restarting the kernel will not rebuild the ACL
        * The document content and access control permission associated with the document will be written to the database when the kernel starts, so when changing the content of published documents or moving published documents after the kernel starts, the published content will not be updated in real-time, the kernel can be restarted to update the database

          * Note: If `Database.Reset = false` is set, restarting the kernel will not update the database
        * When SiYuan kernel is turned off or inaccessible, the published documents can be accessed normally
  3. `Server.Mode.File`: Resource file loading mode, its field value can be set to one of the following three:

      * `dynamic`: Dynamic loading mode, when requesting resource files from the kernel in this mode, the kernel will forward the request to SiYuan kernel in real-time

        * Since SiYuan `assets` directory resources can be accessed publicly when published as long as resource names are known, this option poses some security risks
        * When SiYuan kernel is turned off or inaccessible, resources cannot be loaded
      * `cache`: Dynamic cache mode, when requesting resource files from the kernel in this mode, the kernel first checks whether the corresponding static resource directory (cache directory) contains the file, if not, it requests the file from the SiYuan kernel and saves it to the corresponding static resource directory (cache directory)

        * Since SiYuan `assets` directory resources can be accessed publicly when published as long as resource names are known, this option poses certain security risks
        * Cached resources can be loaded normally when the kernel is closed or cannot be accessed, while uncached resources cannot be loaded
  4. `Server.Index`: Custom site homepage path. Once set, accessing any item in the path list will redirect to the set homepage URL
  5. `Siyuan.Server`: SiYuan kernel service address used to load resources for publishing
  6. `Siyuan.Server.Token`: Access token for SiYuan kernel service. If access authorization code is enabled for the SiYuan kernel service, this option needs to be configured. The token can be copied from <kbd>Settings</kbd> > <kbd>About</kbd> > <kbd>API Token</kbd> in SiYuan

### Start Kernel

* The kernel is an executable program with the filename `publish` located in the root installation directory. If it is run in the same directory as the kernel file, it uses the configuration file `default.config.toml` by default. Otherwise, the resource file path must be specified using the parameter `--config` to start the program without any arguments

### Access Published Content

* Access the homepage

  * `http(s)://host:port`
* Access the document with ID `20200812220555-lj3enxa`

  * `http(s)://host:port/block/20200812220555-lj3enxa`
* Access the document where the block with ID `20210428212840-859h45j` is located

  * `http(s)://host:port/block/20210428212840-859h45j`
* Access the document where the block with ID `20210428212840-859h45j` is located and go directly to that block

  * `http(s)://host:port/block?id=20210428212840-859h45j`
  * `http(s)://host:port/block/20200812220555-lj3enxa?id=20210428212840-859h45j`
* Access the document with ID `20200812220555-lj3enxa` using a light theme

  * `http(s)://host:port/block/20200812220555-lj3enxa?theme=light`
* Access the document with ID `20200812220555-lj3enxa` using a dark theme

  * `http(s)://host:port/block/20200812220555-lj3enxa?theme=dark`
