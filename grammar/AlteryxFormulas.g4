grammar AlteryxFormulas;

formula
    : expr
    ;

expr
    : '(' expr ')'                                                   # parenthesis
    | left=expr '*' right=expr                                       # multiply
    | left=expr '/' right=expr                                       # divide
    | left=expr '+' right=expr                                       # add
    | left=expr '-' right=expr                                       # subtract
    | left=expr '=' right=expr                                       # equal
    | left=expr '>' right=expr                                       # greaterThan
    | left=expr '>=' right=expr                                      # greaterEqual
    | left=expr '<' right=expr                                       # lessThan
    | left=expr '<=' right=expr                                      # lessEqual
    | left=expr '!=' right=expr                                      # notEqual
    | expr 'IN' '(' (expr (',' expr)*)? ')'                          # in
    | expr 'NOT IN' '(' (expr (',' expr)*)? ')'                      # notIn
    | left=expr ('AND'|'&&') right=expr                              # and
    | left=expr ('OR'|'||') right=expr                               # or
    | function                                                       # func
    | 'IF'   expr
      'THEN' expr
      'ELSE' expr
      'ENDIF'                                                        # if
    | 'IF'      expr
      'THEN'    expr
      ('ELSEIF' expr 'THEN' expr)+
      'ELSE' expr
      'ENDIF'                                                        # elseIf
    | Integer                                                        # integer
    | '-'Integer                                                     # integer
    | Decimal                                                        # decimal
    | '-'Decimal                                                     # decimal
    | Datetime                                                       # datetimeLiteral
    | Date                                                           # dateLiteral
    | Field                                                          # field
    | string                                                         # stringLiteral
    ;

string
    : SingleQuoteString
    | DoubleQuoteString
    ;

function
    : Pow '(' expr ',' expr ')'    # pow
    | Min '(' expr (',' expr)+ ')' # min
    | Max '(' expr (',' expr)+ ')' # max
    ;

// Case-insensitive function names
Pow : P O W ;
Min : M I N ;
Max : M A X ;

// Literals
Integer          : Digit+ ;
Decimal          : Digit* '.' Digit+ ;
Date             : ['] DateLiteral [']
                 | '"' DateLiteral '"';
Datetime         : ['] DateTimeLiteral [']
                 | '"' DateTimeLiteral '"';
Field            : '[' ~(']')+ ']' ;
SingleQuoteString: ['] ~(['])* ['] ;
DoubleQuoteString: '"' ~('"')* '"' ;
Whitespace       : [ \t\r\n]+ -> skip ;

// Helpful fragment rules
fragment A:('a'|'A');
fragment B:('b'|'B');
fragment C:('c'|'C');
fragment D:('d'|'D');
fragment E:('e'|'E');
fragment F:('f'|'F');
fragment G:('g'|'G');
fragment H:('h'|'H');
fragment I:('i'|'I');
fragment J:('j'|'J');
fragment K:('k'|'K');
fragment L:('l'|'L');
fragment M:('m'|'M');
fragment N:('n'|'N');
fragment O:('o'|'O');
fragment P:('p'|'P');
fragment Q:('q'|'Q');
fragment R:('r'|'R');
fragment S:('s'|'S');
fragment T:('t'|'T');
fragment U:('u'|'U');
fragment V:('v'|'V');
fragment W:('w'|'W');
fragment X:('x'|'X');
fragment Y:('y'|'Y');
fragment Z:('z'|'Z');
fragment Digit: [0-9];
fragment DateLiteral : Digit Digit Digit Digit '-' Digit Digit '-' Digit Digit ;
fragment DateTimeLiteral: DateLiteral ' ' Digit Digit ':' Digit Digit ':' Digit Digit ;
