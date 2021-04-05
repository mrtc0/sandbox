import { HtmlSanitizer } from './sanitizer'

export const createSanitizer = () => {
  return new HtmlSanitizer()
}

const sanitizer = createSanitizer()
sanitizer.sanitize("<img src='x' onerror=alert(1)>")