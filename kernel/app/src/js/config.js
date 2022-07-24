/* 覆盖 *.config.toml 配置文件中定义的字段 */
// window.siyuan = {
//     config: {
//         appearance: { mode: 0, codeBlockThemeDark: "atom-one-dark", codeBlockThemeLight: "atom-one-light" },
//         editor: {
//             fontSize: 16,
//             codeLineWrap: true,
//             codeLigatures: true,
//             codeSyntaxHighlightLineNum: true,
//             plantUMLServePath: "https://www.plantuml.com/plantuml/svg/~1",
//         },
//     },
//     languages: { copy: "复制" },
// };


/**
 * 设置编辑器字号
 * REF https://github.com/siyuan-note/siyuan/blob/fcabf93cabf0383a8b59616d66ec44e7869236cf/app/src/protyle/export/index.ts#L242-L107
 * @params {number} fontSize 字号
 * @return {number} 设置后的字号
 * @return {null} 没有找到字号
 */
window.publish.setFontSize = (fontSize, fontFamily, codeLigatures) => {
    let style = document.getElementById('editorFontSize');
    if (style) {
        const height = Math.floor(fontSize * 1.625);
        style.innerHTML = `
.b3-typography, .protyle-wysiwyg, .protyle-title {font-size:${fontSize}px !important}
.b3-typography code:not(.hljs), .protyle-wysiwyg code:not(.hljs) { font-variant-ligatures: ${codeLigatures ? "normal" : "none"} }
.li > .protyle-action {height:${height + 8}px;line-height: ${height + 8}px}
.protyle-wysiwyg [data-node-id].li > .protyle-action ~ .h1, .protyle-wysiwyg [data-node-id].li > .protyle-action ~ .h2, .protyle-wysiwyg [data-node-id].li > .protyle-action ~ .h3, .protyle-wysiwyg [data-node-id].li > .protyle-action ~ .h4, .protyle-wysiwyg [data-node-id].li > .protyle-action ~ .h5, .protyle-wysiwyg [data-node-id].li > .protyle-action ~ .h6 {line-height:${height + 8}px;}
.protyle-wysiwyg [data-node-id].li > .protyle-action:after {height: ${fontSize}px;width: ${fontSize}px;margin:-${fontSize / 2}px 0 0 -${fontSize / 2}px}
.protyle-wysiwyg [data-node-id].li > .protyle-action svg {height: ${Math.max(14, fontSize - 8)}px}
.protyle-wysiwyg [data-node-id] [spellcheck="false"] {min-height:${height}px}
.protyle-wysiwyg .li {min-height:${height + 8}px}
.protyle-gutters button svg {height:${height}px}
.protyle-wysiwyg img.emoji, .b3-typography img.emoji {width:${height - 8}px}
.protyle-wysiwyg .h1 img.emoji, .b3-typography h1 img.emoji {width:${Math.floor(fontSize * 1.75 * 1.25)}px}
.protyle-wysiwyg .h2 img.emoji, .b3-typography h2 img.emoji {width:${Math.floor(fontSize * 1.55 * 1.25)}px}
.protyle-wysiwyg .h3 img.emoji, .b3-typography h3 img.emoji {width:${Math.floor(fontSize * 1.38 * 1.25)}px}
.protyle-wysiwyg .h4 img.emoji, .b3-typography h4 img.emoji {width:${Math.floor(fontSize * 1.25 * 1.25)}px}
.protyle-wysiwyg .h5 img.emoji, .b3-typography h5 img.emoji {width:${Math.floor(fontSize * 1.13 * 1.25)}px}
.protyle-wysiwyg .h6 img.emoji, .b3-typography h6 img.emoji {width:${Math.floor(fontSize * 1.25)}px}
.b3-typography, .protyle-wysiwyg, .protyle-title, .protyle-title__input{font-family: ${fontFamily} "quote", "Helvetica Neue", "Luxi Sans", "DejaVu Sans", "Hiragino Sans GB", "Microsoft Yahei", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Noto Color Emoji", "Segoe UI Symbol", "Android Emoji", "EmojiSymbols" !important;}
`;
        return fontSize;
    }
    return null;
}

(() => {
    const url = new URL(window.location.href);

    /* 设置字体样式 */
    window.publish.setFontSize(
        window.siyuan.config.editor.fontSize,
        window.siyuan.config.editor.fontFamily,
        window.siyuan.config.editor.codeLigatures,
    );

    /* 设置题头 */
    document.getElementById('title').style.minHeight = window.publish.config.title.height;
    const title_img = window.publish.config.title.img
    if (title_img) {
        const img = document.getElementById('title-img');
        const reg = /background-image:url\((.*)\);/
        if (reg.test(title_img)) { // 使用了自定义背景图片
            const src = reg.exec(title_img)[1];
            const style = title_img.replace(/background-image:url\((.*)\);/, "");
            img.src = src;
            img.style = style;
        } else { // 使用了随机背景图片
            img.style = title_img;
        }
    }

    /* 设置主题模式 */
    function setThemeMode(mode) {
        switch (mode) {
            case 0:
                document.getElementById('themeDefaultStyle').href = window.publish.config.theme.light.default;
                document.getElementById('themeStyle').href = window.publish.config.theme.light.theme;
                document.getElementById('themeCustomStyle').href = window.publish.config.theme.light.custom;
                window.siyuan.config.appearance.mode = 0;
                break;
            case 1:
                document.getElementById('themeDefaultStyle').href = window.publish.config.theme.dark.default;
                document.getElementById('themeStyle').href = window.publish.config.theme.dark.theme;
                document.getElementById('themeCustomStyle').href = window.publish.config.theme.dark.custom;
                window.siyuan.config.appearance.mode = 1;
                break;
        }
    }

    /* 使用 URL 参数 theme 设置主题模式 */
    switch (url.searchParams.get('theme')) {
        case 'light':
            setThemeMode(0);
            break;
        case 'dark':
            setThemeMode(1);
            break;
        case 'auto':
        default:
            /* 使用服务器配置设置主题模式 */
            switch (window.siyuan.config.appearance.mode) {
                case 0:
                default:
                    setThemeMode(0);
                    break;
                case 1:
                    setThemeMode(1);
                    break;
                case 2:
                    /* 使用浏览器配置设置主题模式 */
                    switch (true) {
                        case window.matchMedia('(prefers-color-scheme: light)').matches:
                            setThemeMode(0);
                            break;
                        case window.matchMedia('(prefers-color-scheme: dark)').matches:
                            setThemeMode(1);
                            break;
                        default:
                            setThemeMode(0);
                            break;
                    }
                    break;
            }
    }
})()
