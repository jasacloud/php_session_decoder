package php_session_decoder

import "github.com/jasacloud/php_session_decoder/php_serialize"

const SEPARATOR_VALUE_NAME rune = '|'

type PhpSession map[string]php_serialize.PhpValue
