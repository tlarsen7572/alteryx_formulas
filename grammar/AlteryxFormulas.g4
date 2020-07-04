grammar AlteryxFormulas;

formula
    : fieldExpr                                                      # formulaIsField
    | numberExpr                                                     # formulaIsNumber
    | dateExpr                                                       # formulaIsDate
    | stringExpr                                                     # formulaIsString
    | boolExpr                                                       # formulaIsBool
    ;

fieldExpr
    : '(' Field ')'                                                  # fieldParenthesis
    | Field                                                          # anyField
    ;

stringExpr
    : '(' stringExpr ')'                                             # stringParenthesis
    | left=stringExpr '+' right=stringExpr                           # concatenate
    | If   boolExpr
      Then stringExpr
      Else stringExpr
      Endif                                                          # stringIf
    | If      boolExpr
      Then    stringExpr
      (Elseif boolExpr Then stringExpr)+
      Else    stringExpr
      Endif                                                          # stringElseIf
    | stringFunction                                                 # stringFunc
    | str                                                            # stringLiteral
    | Field                                                          # stringField
    ;

stringFunction
    : Min '(' stringExpr (',' stringExpr)+ ')'                          # stringMin
    | Max '(' stringExpr (',' stringExpr)+ ')'                          # stringMax
    ;

numberExpr
    : '(' numberExpr ')'                                             # numberParenthesis
    | left=numberExpr '*' right=numberExpr                           # multiply
    | left=numberExpr '/' right=numberExpr                           # divide
    | left=numberExpr '+' right=numberExpr                           # add
    | left=numberExpr '-' right=numberExpr                           # subtract
    | numberFunction                                                 # numberFunc
    | If   boolExpr
      Then numberExpr
      Else numberExpr
      Endif                                                          # numberIf
    | If      boolExpr
      Then    numberExpr
      (Elseif boolExpr Then numberExpr)+
      Else    numberExpr
      Endif                                                          # numberElseIf
    | Integer                                                        # integer
    | '-'Integer                                                     # integer
    | Decimal                                                        # decimal
    | '-'Decimal                                                     # decimal
    | Field                                                          # numberField
    ;

numberFunction
    : Pow '(' numberExpr ',' numberExpr ')'                          # pow
    | Min '(' numberExpr (',' numberExpr)+ ')'                       # numberMin
    | Max '(' numberExpr (',' numberExpr)+ ')'                       # numberMax
    | Null '()'                                                      # numberNull
    ;

dateExpr
    : '(' dateExpr ')'                                               # dateParenthesis
    | If   boolExpr
      Then dateExpr
      Else dateExpr
      Endif                                                          # dateIf
    | If      boolExpr
      Then    dateExpr
      (Elseif boolExpr Then dateExpr)+
      Else    dateExpr
      Endif                                                          # dateElseIf
    | dateFunction                                                   # dateFunc
    | Datetime                                                       # datetimeLiteral
    | Date                                                           # dateLiteral
    | Field                                                          # dateField
    ;

dateFunction
    : Min '(' dateExpr (',' dateExpr)+ ')'                           # dateMin
    | Max '(' dateExpr (',' dateExpr)+ ')'                           # dateMax
    ;

boolExpr
    : '(' boolExpr ')'                                               # boolParenthesis
    | left=stringExpr '=' right=stringExpr                           # stringEqual
    | left=stringExpr '>' right=stringExpr                           # stringGreaterThan
    | left=stringExpr '>=' right=stringExpr                          # stringGreaterEqual
    | left=stringExpr '<' right=stringExpr                           # stringLessThan
    | left=stringExpr '<=' right=stringExpr                          # stringLessEqual
    | left=stringExpr '!=' right=stringExpr                          # stringNotEqual
    | stringExpr In '(' (stringExpr (',' stringExpr)*)? ')'          # stringIn
    | stringExpr Not In '(' (stringExpr (',' stringExpr)*)? ')'      # stringNotIn
    | left=numberExpr '=' right=numberExpr                           # numberEqual
    | left=numberExpr '>' right=numberExpr                           # numberGreaterThan
    | left=numberExpr '>=' right=numberExpr                          # numberGreaterEqual
    | left=numberExpr '<' right=numberExpr                           # numberLessThan
    | left=numberExpr '<=' right=numberExpr                          # numberLessEqual
    | left=numberExpr '!=' right=numberExpr                          # numberNotEqual
    | numberExpr In '(' (numberExpr (',' numberExpr)*)? ')'          # numberIn
    | numberExpr Not In '(' (numberExpr (',' numberExpr)*)? ')'      # numberNotIn
    | left=dateExpr '=' right=dateExpr                               # dateEqual
    | left=dateExpr '>' right=dateExpr                               # dateGreaterThan
    | left=dateExpr '>=' right=dateExpr                              # dateGreaterEqual
    | left=dateExpr '<' right=dateExpr                               # dateLessThan
    | left=dateExpr '<=' right=dateExpr                              # dateLessEqual
    | left=dateExpr '!=' right=dateExpr                              # dateNotEqual
    | dateExpr In '(' (dateExpr (',' dateExpr)*)? ')'                # dateIn
    | dateExpr Not In '(' (dateExpr (',' dateExpr)*)? ')'            # dateNotIn
    | left=boolExpr (And | '&&') right=boolExpr                      # and
    | left=boolExpr (Or | '||') right=boolExpr                       # or
    | Bool                                                           # boolLiteral
    | Field                                                          # boolField
    ;

str
    : SingleQuoteString
    | DoubleQuoteString
    ;

// Case-insensitive function names
Pow : P O W ;
Min : M I N ;
Max : M A X ;
Null: N U L L ;

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
