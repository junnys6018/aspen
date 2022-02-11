package main

type Parser struct {
	tokens  TokenStream
	current int
}

// Grammar

func (p *Parser) expression() Expression {
	return p.logicOr()
}

func (p *Parser) logicOr() Expression {
	expr := p.logicAnd()

	for p.match(TOKEN_PIPE_PIPE) {
		operator := p.previous()
		right := p.logicAnd()
		expr = &BinaryExpression{left: expr, right: right, operator: *operator}
	}

	return expr
}

func (p *Parser) logicAnd() Expression {
	expr := p.equality()

	for p.match(TOKEN_AMP_AMP) {
		operator := p.previous()
		right := p.equality()
		expr = &BinaryExpression{left: expr, right: right, operator: *operator}
	}

	return expr
}

func (p *Parser) equality() Expression {
	expr := p.comparison()

	for p.match(TOKEN_EQUAL_EQUAL, TOKEN_BANG_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = &BinaryExpression{left: expr, right: right, operator: *operator}
	}

	return expr
}

func (p *Parser) comparison() Expression {
	expr := p.bitOr()

	for p.match(TOKEN_GREATER, TOKEN_GREATER_EQUAL, TOKEN_LESS, TOKEN_LESS_EQUAL) {
		operator := p.previous()
		right := p.bitOr()
		expr = &BinaryExpression{left: expr, right: right, operator: *operator}
	}

	return expr
}

func (p *Parser) bitOr() Expression {
	expr := p.bitXor()

	for p.match(TOKEN_PIPE) {
		operator := p.previous()
		right := p.bitXor()
		expr = &BinaryExpression{left: expr, right: right, operator: *operator}
	}

	return expr
}

func (p *Parser) bitXor() Expression {
	expr := p.bitAnd()

	for p.match(TOKEN_CARET) {
		operator := p.previous()
		right := p.bitAnd()
		expr = &BinaryExpression{left: expr, right: right, operator: *operator}
	}

	return expr
}

func (p *Parser) bitAnd() Expression {
	expr := p.term()

	for p.match(TOKEN_AMP) {
		operator := p.previous()
		right := p.term()
		expr = &BinaryExpression{left: expr, right: right, operator: *operator}
	}

	return expr
}

func (p *Parser) term() Expression {
	expr := p.factor()

	for p.match(TOKEN_MINUS, TOKEN_PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = &BinaryExpression{left: expr, right: right, operator: *operator}
	}

	return expr
}

func (p *Parser) factor() Expression {
	expr := p.unary()

	for p.match(TOKEN_SLASH, TOKEN_STAR) {
		operator := p.previous()
		right := p.unary()
		expr = &BinaryExpression{left: expr, right: right, operator: *operator}
	}

	return expr
}

func (p *Parser) unary() Expression {
	if p.match(TOKEN_BANG, TOKEN_MINUS) {
		operator := p.previous()
		right := p.unary()
		return &UnaryExpression{operand: right, operator: *operator}
	}

	return p.primary()
}

func (p *Parser) primary() Expression {
	if p.match(TOKEN_FALSE, TOKEN_TRUE, TOKEN_NIL, TOKEN_INT, TOKEN_FLOAT, TOKEN_STRING) {
		return &LiteralExpression{value: *p.previous()}
	}

	if p.match(TOKEN_LEFT_PAREN) {
		expr := p.expression()
		p.consume(TOKEN_RIGHT_PAREN)
		return &GroupingExpression{expr: expr}
	}

	return nil
}

// Helpers

func (p *Parser) consume(tokenType TokenType) *Token {
	if p.check(tokenType) {
		return p.advance()
	}

	return nil
}

func (p *Parser) match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().tokenType == tokenType
}

func (p *Parser) advance() *Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().tokenType == TOKEN_EOF
}

func (p *Parser) peek() *Token {
	return &p.tokens[p.current]
}

func (p *Parser) previous() *Token {
	return &p.tokens[p.current-1]
}

func Parse(tokens TokenStream) (Expression, error) {
	parser := Parser{tokens: tokens, current: 0}
	expr := parser.expression()
	return expr, nil
}