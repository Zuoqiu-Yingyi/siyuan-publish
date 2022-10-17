/* 新窗口打开 */
export function open(url, terget, left, top, width, height) {
    window.top.open(
        url,
        terget,
        `
            popup = true,
            left = ${left},
            top = ${top},
            width = ${width},
            height = ${height},
        `,
    )
}
