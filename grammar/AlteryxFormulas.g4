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
    | CharFromInt '(' expr ')'                                       # charFromIntFunc
    | CharToInt '(' expr ')'                                         # charToIntFunc
    | Contains '(' expr ',' expr (',' expr)? ')'                     # containsFunc
    | Cos '(' expr ')'                                               # cosFunc
    | Cosh '(' expr ')'                                              # coshFunc
    | CountWords '(' expr ')'                                        # countWordsFunc
    | Distance '(' expr ',' expr ',' expr ',' expr ')'               # distanceFunc
    | EndsWith '(' expr ',' expr (',' expr)? ')'                     # endsWithFunc
    | Exp '(' expr ')'                                               # expFunc
    | FindString '(' expr ',' expr ')'                               # findStringFunc
    | Floor '(' expr ')'                                             # floorFunc
    | GetWord '(' expr ',' expr ')'                                  # getWordFunc
    | HexToNumber '(' expr ')'                                       # hexToNumberFunc
    | Left '(' expr ',' expr ')'                                     # leftFunc
    | Length '(' expr ')'                                            # lengthFunc
    | Log '(' expr ')'                                               # logFunc
    | Log10 '(' expr ')'                                             # log10Func
    | Lowercase '(' expr ')'                                         # lowercaseFunc
    | Max '(' expr (',' expr)+ ')'                                   # maxFunc
    | Median '(' expr (',' expr)* ')'                                # medianFunc
    | Min '(' expr (',' expr)+ ')'                                   # minFunc
    | Mod '(' expr (',' expr)* ')'                                   # modFunc
    | Null '()'                                                      # nullFunc
    | PadLeft '(' expr ',' expr ',' expr ')'                         # padLeftFunc
    | Pi '()'                                                        # piFunc
    | Pow '(' expr ',' expr ')'                                      # powFunc
    | Rand '()'                                                      # randFunc
    | RandInt '(' expr ')'                                           # randIntFunc
    | Round '(' expr ',' expr ')'                                    # roundFunc
    | Sin '(' expr ')'                                               # sinFunc
    | Sinh '(' expr ')'                                              # sinhFunc
    | Sqrt '(' expr ')'                                              # sqrtFunc
    | Switch '(' expr ',' expr (',' expr ',' expr)+ ')'              # switchFunc
    | Tan '(' expr ')'                                               # tanFunc
    | Tanh '(' expr ')'                                              # tanhFunc
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
Abs         : A B S ;
Acos        : A C O S ;
Asin        : A S I N ;
Atan        : A T A N ;
Atan2       : A T A N '2' ;
Average     : A V E R A G E ;
Ceil        : C E I L ;
CharFromInt : C H A R F R O M I N T ;
CharToInt   : C H A R T O I N T ;
Contains    : C O N T A I N S ;
Cos         : C O S ;
Cosh        : C O S H ;
CountWords  : C O U N T W O R D S ;
Distance    : D I S T A N C E ;
EndsWith    : E N D S W I T H ;
FindString  : F I N D S T R I N G ;
Exp         : E X P ;
Floor       : F L O O R ;
GetWord     : G E T W O R D ;
HexToNumber : H E X T O N U M B E R ;
Left        : L E F T ;
Length      : L E N G T H ;
Log         : L O G ;
Log10       : L O G '10' ;
Lowercase   : L O W E R C A S E ;
Max         : M A X ;
Median      : M E D I A N ;
Min         : M I N ;
Mod         : M O D ;
Null        : N U L L ;
PadLeft     : P A D L E F T ;
Pi          : P I ;
Pow         : P O W ;
Rand        : R A N D ;
RandInt     : R A N D I N T ;
Round       : R O U N D ;
Sin         : S I N ;
Sinh        : S I N H ;
Sqrt        : S Q R T ;
Switch      : S W I T C H ;
Tan         : T A N ;
Tanh        : T A N H ;
Iif         : I I F ;

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
