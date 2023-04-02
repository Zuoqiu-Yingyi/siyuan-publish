<h1>English Documentation</h1>
This document was translated by ChatGPT-4
<h2>Quick Start</h2>
<h3>Requirements</h3>
<ul>
<li>The device has installed Siyuan Note or can access Siyuan Note server service</li>
<li>Download the latest app distribution for the device <code>Operating System</code> and <code>Architecture</code> in <a href="https://github.com/Zuoqiu-Yingyi/siyuan-publish/releases">Releases</a></li>
<li>Unless otherwise stated, <code>kernel</code> below refers to the executable program of this application, and <code>Siyuan kernel</code> refers to the application program providing Siyuan Note server functionality</li>
</ul>
<h3>Installation</h3>
<ul>
<li>Move the latest app installation package <code>*.zip</code> for the device <code>Operating System</code> and <code>Architecture</code> to the installation directory and unzip it</li>
</ul>
<h3>Project Structure</h3>

<div style="position: relative">
	<pre><code class="hljs language-plaintext">.
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
      ├─appearance // Appearance resource directory (corresponding to Siyuan Note workspace /conf/appearance/ directory)
      ├─aeests // Document resource file directory (corresponding to Siyuan Note workspace /data/assets/ directory)
      ├─emojis // Emotion resource file directory (corresponding to Siyuan Note workspace /data/emojis/ directory)
      ├─export // Export resource file directory (corresponding to Siyuan Note workspace /temp/export/ directory)
      ├─stage // Rendering tool resource file directory (corresponding to Siyuan Note installation /resource/stage/ directory)
      └─widgets // Widget resource file directory (corresponding to Siyuan Note workspace /data/widgets/ directory)
</code></pre>

<ul>
<li>
<p>The configuration file format is TOML, for more information, please refer to <a href="https://toml.io/cn/">TOML: A human-friendly configuration file format</a></p>
<ul>
<li>Full specification: <a href="https://toml.io/cn/latest">TOML v1.0.0</a></li>
</ul>
</li>
<li>
<p>Before running the kernel for the first time, you need to set the following fields:</p>
<ol>
<li>
<p><code>Server.Port</code>: Web service access port, the default is port <code>80</code></p>
</li>
<li>
<p><code>Server.Mode.Page</code>: Document page loading mode, field value can be set to one of the following three:</p>
<ul>
<li>
<p><code>dynamic</code>: Dynamic loading mode, when set to this mode, the kernel will request data from the Siyuan kernel in real-time</p>
<ul>
<li>The kernel determines in real-time whether the accessed document is public</li>
<li>The kernel does not use the database cache Access Control List (ACL), so changing the access control permission of a document by editing the custom attribute&lt;kbd&gt;publish-access&lt;/kbd&gt; can take effect in real-time</li>
<li>Siyuan kernel cannot be turned off, the published content cannot be accessed normally when the Siyuan kernel is turned off or the Siyuan kernel service is inaccessible</li>
</ul>
</li>
<li>
<p><code>cache</code>: Dynamic cache mode, when set to this mode, the kernel will first query the cache content from the database, and if the cache content is not found, it will request data from the Siyuan kernel and write it to the cache</p>
<ul>
<li>
<p>The kernel will request data from the Siyuan kernel to establish an ACL at startup and write the ACL to the database, so when editing the custom attribute&lt;kbd&gt;publish-access&lt;/kbd&gt; to change the access control permission of a document, it cannot take effect in real-time and must be restarted in the kernel to update the ACL</p>
<ul>
<li>Note: If <code>Database.Reset = false</code> is set, restarting the kernel will not rebuild the ACL</li>
</ul>
</li>
<li>
<p>The document content and the access control permission associated with the document will be written to the database after the first access, so after changing the cached document content or moving the cached document, the published content will not be updated in real-time, and the kernel needs to be restarted to clear the cache content</p>
<ul>
<li>Note: If <code>Database.Reset = false</code> is set, restarting the kernel will not clear the cache</li>
</ul>
</li>
<li>
<p>Siyuan kernel cannot be turned off, the cached content can be accessed normally when the Siyuan kernel is turned off or the Siyuan kernel service is inaccessible, and the content outside the cache cannot be accessed normally</p>
</li>
</ul>
</li>
<li>
<p><code>static</code>: Static loading mode, when set to this mode, the kernel will write all published document information to the database at startup</p>
<ul>
<li>
<p>The kernel will request data from the Siyuan kernel to establish an ACL at startup and write the ACL to the database, so when editing the custom attribute&lt;kbd&gt;publish-access&lt;/kbd&gt; to change the access control permission of a document, it cannot take effect in real-time and must be restarted in the kernel to update the ACL</p>
<ul>
<li>Note: If <code>Database.Reset = false</code> is set, restarting the kernel will not rebuild the ACL</li>
</ul>
</li>
<li>
<p>The document content and access control permission associated with the document will be written to the database when the kernel starts, so when changing the content of published documents or moving published documents after the kernel starts, the published content will not be updated in real-time, the kernel can be restarted to update the database</p>
<ul>
<li>Note: If <code>Database.Reset = false</code> is set, restarting the kernel will not update the database</li>
</ul>
</li>
<li>
<p>When Siyuan kernel is turned off or inaccessible, the published documents can be accessed normally</p>
</li>
</ul>
</li>
</ul>
</li>
<li>
<p><code>Server.Mode.File</code>: Resource file loading mode, its field value can be set to one of the following three:</p>
<ul>
<li>
<p><code>dynamic</code>: Dynamic loading mode, when requesting resource files from the kernel in this mode, the kernel will forward the request to Siyuan kernel in real-time</p>
<ul>
<li>Since Siyuan <code>assets</code> directory resources can be accessed publicly when published as long as resource names are known, this option poses some security risks</li>
<li>When Siyuan kernel is turned off or inaccessible, resources cannot be loaded</li>
</ul>
</li>
<li>
<p><code>cache</code>: Dynamic cache mode, when requesting resource files from the kernel in this mode, the kernel first checks whether the corresponding static resource directory (cache directory) contains the file, if not, it requests the file from the Siyuan kernel and saves it to the corresponding static resource directory (cache directory)</p>
<ul>
<li>Since Siyuan <code>assets</code> directory resources can be accessed publicly when published as long as resource names are known,</li>
</ul>
</li>
</ul>
</li>
</ol>
</li>
</ul>
