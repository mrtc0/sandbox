export interface HtmlSanitizer {
  sanitize(html: string): string
  createTreeWalker(document: Document, node: HTMLDivElement): any
  shouldRemoveNode(node: Node): boolean
}

export class HtmlSanitizer implements HtmlSanitizer {
  sanitize(html: string) {
    const sandbox = document.implementation.createHTMLDocument('')
    const root = sandbox.createElement('div')
    root.innerHTML = html

    var treeWalker = this.createTreeWalker(sandbox, root)
    var node = treeWalker.firstChild()

    if (!node) { return "" }

    do {
      if (node.nodeType === Node.TEXT_NODE) {
        console.log("text node")
      }

      if (node.nodeType === Node.COMMENT_NODE) {
        console.log("comment node")
      }

      if (this.shouldRemoveNode(node)) {
        console.log(`remove node ${node}`)
        root.removeChild(node)
      }

    } while ((node = treeWalker.nextSibling()))

    console.log(root)
    return root
  }

  shouldRemoveNode(node: Node): boolean {
    return node.nodeName.toLowerCase() === 'img'
  }

  createTreeWalker(document: Document, node: HTMLDivElement) {
    return document.createTreeWalker(node, NodeFilter.SHOW_TEXT | NodeFilter.SHOW_ELEMENT | NodeFilter.SHOW_COMMENT, null, false);
  }
}
