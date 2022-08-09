import { Base } from './../base/index.js';

export {
    Font as Plugin,
};

class Font extends Base {
    static META = {
        NAME: 'publish-font',
        UUID: 'C5F604F5-FE69-492C-8997-82145845CDA7',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: '字体设置(字体族, 字号)',
        BEFORE: {
            async: true,
            defer: false,
        },
    };

    constructor(context) {
        super(context);
    }

    async before() {
        /* 设置字体配置 */
        this.setFontSize(
            window.siyuan.config.editor.fontSize,
            window.siyuan.config.editor.fontFamily,
            window.siyuan.config.editor.codeLigatures,
        );
    }

    /**
     * 设置编辑器字号
     * REF https://github.com/siyuan-note/siyuan/blob/fcabf93cabf0383a8b59616d66ec44e7869236cf/app/src/protyle/export/index.ts#L242-L107
     * @params {number} fontSize 字号
     * @return {number} 设置后的字号
     * @return {null} 没有找到字号
     */
    setFontSize(fontSize, fontFamily, codeLigatures) {
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
}
