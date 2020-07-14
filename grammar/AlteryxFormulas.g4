grammar AlteryxFormulas;

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
    | expr In '(' (expr (',' expr)*)? ')'                            # in
    | expr Not In '(' (expr (',' expr)*)? ')'                        # notIn
    | left=expr (And | '&&') right=expr                              # and
    | left=expr (Or | '||') right=expr                               # or
    | If      ifStmnt=expr
      Then    thenStmnt=expr
      (Elseif ifStmnt=expr Then thenStmnt=expr)*
      Else    elseStmnt=expr
      Endif                                                          # exprIf
    | Iif '(' expr ',' expr ',' expr ')'                             # iifFunc
    | Abs '(' expr ')'                                               # absFunc
    | Acos '(' expr ')'                                              # acosFunc
    | Asin '(' expr ')'                                              # asinFunc
    | Atan '(' expr ')'                                              # atanFunc
    | Atan2 '(' expr ',' expr ')'                                    # atan2Func
    | Average '(' expr (',' expr)* ')'                               # averageFunc
    | Ceil '(' expr ')'                                              # ceilFunc
    | Cos '(' expr ')'                                               # cosFunc
    | Cosh '(' expr ')'                                              # coshFunc
    | Distance '(' expr ',' expr ',' expr ',' expr ')'               # distanceFunc
    | Null '()'                                                      # nullFunc
    | Pow '(' expr ',' expr ')'                                      # powFunc
    | Min '(' expr (',' expr)+ ')'                                   # minFunc
    | Max '(' expr (',' expr)+ ')'                                   # maxFunc
    | Integer                                                        # numberLiteral
    | '-'Integer                                                     # numberLiteral
    | Decimal                                                        # numberLiteral
    | '-'Decimal                                                     # numberLiteral
    | SingleQuoteString                                              # stringLiteral
    | DoubleQuoteString                                              # stringLiteral
    | Datetime                                                       # datetimeLiteral
    | Date                                                           # dateLiteral
    | Bool                                                           # boolLiteral
    | Field                                                          # exprField
    ;

// Case-insensitive function names
Abs     : A B S ;
Acos    : A C O S ;
Asin    : A S I N ;
Atan    : A T A N ;
Atan2   : A T A N '2' ;
Average : A V E R A G E ;
Ceil    : C E I L ;
Cos     : C O S ;
Cosh    : C O S H ;
Distance: D I S T A N C E ;
Pow     : P O W ;
Min     : M I N ;
Max     : M A X ;
Null    : N U L L ;
Iif     : I I F ;

// Case-insensitive keywords
In     : I N ;
Not    : N O T ;
And    : A N D ;
Or     : O R ;
If     : I F ;
Then   : T H E N ;
Else   : E L S E ;
Elseif : E L S E I F ;
Endif  : E N D I F ;

// Literals
Bool             : TrueLiteral | FalseLiteral ;
Integer          : Digit+ ;
Decimal          : Digit* '.' Digit+ ;
Date             : [`] DateLiteral [`] ;
Datetime         : [`] DateTimeLiteral [`] ;
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
fragment TrueLiteral : T R U E ;
fragment FalseLiteral : F A L S E ;
