import { Base } from './../base/index.js';

export {
    Icon as Plugin,
};

class Icon extends Base {
    static META = {
        NAME: 'publish-icon',
        UUID: 'CE409332-BC12-41A8-AC7C-087E49B21A56',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: '图标处理',
        DEPENDENCY: [],
    };

    constructor(context) {
        super(context);
        this.TYPE_ICON_MAP = {
            NodeAudio: "#iconRecord",
            NodeBlockQueryEmbed: "#iconSQL",
            NodeBlockquote: "#iconQuote",
            NodeCodeBlock: "#iconCode",
            NodeDocument: "#iconFile",
            NodeHTMLBlock: "#iconHTML5",
            NodeHeading: {
                h1: "#iconH1",
                h2: "#iconH2",
                h3: "#iconH3",
                h4: "#iconH4",
                h5: "#iconH5",
                h6: "#iconH6",
            },
            NodeIFrame: "#iconLanguage",
            NodeList: {
                o: "#iconList",
                u: "#iconOrderedList",
                t: "#iconCheck",
            },
            NodeListItem: "#iconListItem",
            NodeMathBlock: "#iconMath",
            NodeParagraph: "#iconParagraph",
            NodeSuperBlock: "#iconSuper",
            NodeTable: "#iconTable",
            NodeThematicBreak: "#iconLine",
            NodeVideo: "#iconVideo",
            NodeWidget: "#iconBoth",
        }

        this.context.meta.set('TYPE_ICON_MAP', this.TYPE_ICON_MAP);
    }

    /**
     * 创建图标
     * @params {string} icon: 图标名
     * @params {string} className: 类名
     * @params {string} id: 块 ID
     * @returns {HTMLElement} 图标
     */
    createIcon(icon, className, id = null) {
        const element = this.context.document.createElement('div');
        element.innerHTML = `
        <svg class="${className}">
            <use xlink:href="${icon}"></use>
        </svg>
    `;
        if (id) element.firstElementChild.dataset.nodeId = id;
        return element.firstElementChild;
    }
}
