/*
	A SWITCH STATEMENT IS A MULTIWAY BRANCH STATEMENT.
	IT PROVIDES AN EFICIENT WAY TO
	TRANSFER THE EXECUTION TO DIFFERENT PARTS OF A CODE BASED ON THE
	VALUE ( ALSO CALLED CASES ) OF THE EXPRESSION. GO LANGUAGE SUPPORTS
	TWO TYPES OF SWITCH STATEMENTS:-

	1.Expression Switch
	2. Type Switch

					EXPRESSION SWITCH
					-----------------

	EXPRESSION SWITCH IS SIMILAR STATEMENT IN C,C++ OR JAVA, IT PROVIDES
	AN EASY WAY TO DISPATCH EXECUTION TO DIFFERENT PARTS OF CODE BASED ON
	THE VALUE OF THE EXPRESSION
*/

/*
 			switch optstatement; optexpression{
			case expression1: Statement..
			case expression2: Statement..
			...
			default: Statement..
}
*/
package main

func Switche() {

	/*
		IMPORTANT POINTS:-

		1. BOTH optstatement AND optexprssion IN THE EXPRESSION SWITCH
			ARE OPTIONAL STATEMENTS.
		2. IF BOTH optstatements AND optexpression ARE PRESENT, THEN A
			SEMI-COLON(;)IS REQUIRED IN BETWEEN THEM.
		3. IF THE SWITCH STATEMENT DOES NOT CONTAIN ANY EXPRESSION THEN
			THE COMPILER AASSUMES THAT THE EXPRESSION IS TRUE.
		4. THE OPTIONAL STATEMENTS, i.e, optstatements CONTAINS SIMPLE
			STATEMENT LIKE VARIABLE DECLARATIONS, INCREMENT ASSIGNMENT OR FUNCTION CALL ETC
		5. IF A VARIABLE PRESENT IN THE OPTIONAL STATEMENTS, THEN THE SCOPE OF THE VARAIABLE IS LIMITED TO THAT
			SWITCH STATEMENT.
		6. In switch statement, the case and default statement does not contain any break statement.
			But you are allowed to use break and fallthrough statement if your program required.

		7. The default statement is optional in switch statement.
		8. If a case can contain multiple values and these values are separated by comma(,).
		9. If a case does not contain any expression, then the compiler assume that te expression is true.

	*/

}
