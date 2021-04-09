{{ define "*ast.ArrayType" -}}
{{ template "prefix1" }}ArrayType
{{ template "prefix2" }}├── Lbrack = {{ position .Lbrack}}
{{ template "prefix2" }}├── Len
{{with .Len}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Elt
{{with .Elt}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.AssignStmt" -}}
{{ template "prefix1" }}AssignStmt
{{ template "prefix2" }}├── Lhs = (length={{ len .Lhs }})
{{ range .Lhs }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── TokPos = {{ position .TokPos}}
{{ template "prefix2" }}├── Tok = {{ .Tok}}
{{ template "prefix2" }}└── Rhs = (length={{ len .Rhs }})
{{ range .Rhs }}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.BadDecl" -}}
{{ template "prefix1" }}BadDecl
{{ template "prefix2" }}├── From = {{ position .From}}
{{ template "prefix2" }}└── To = {{ position .To}}
{{ end }}
{{ define "*ast.BadExpr" -}}
{{ template "prefix1" }}BadExpr
{{ template "prefix2" }}├── From = {{ position .From}}
{{ template "prefix2" }}└── To = {{ position .To}}
{{ end }}
{{ define "*ast.BadStmt" -}}
{{ template "prefix1" }}BadStmt
{{ template "prefix2" }}├── From = {{ position .From}}
{{ template "prefix2" }}└── To = {{ position .To}}
{{ end }}
{{ define "*ast.BasicLit" -}}
{{ template "prefix1" }}BasicLit
{{ template "prefix2" }}├── ValuePos = {{ position .ValuePos}}
{{ template "prefix2" }}├── Kind = {{ .Kind}}
{{ template "prefix2" }}└── Value = {{ .Value}}
{{ end }}
{{ define "*ast.BinaryExpr" -}}
{{ template "prefix1" }}BinaryExpr
{{ template "prefix2" }}├── X
{{with .X}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── OpPos = {{ position .OpPos}}
{{ template "prefix2" }}├── Op = {{ .Op}}
{{ template "prefix2" }}└── Y
{{with .Y}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.BlockStmt" -}}
{{ template "prefix1" }}BlockStmt
{{ template "prefix2" }}├── Lbrace = {{ position .Lbrace}}
{{ template "prefix2" }}├── List = (length={{ len .List }})
{{ range .List }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Rbrace = {{ position .Rbrace}}
{{ end }}
{{ define "*ast.BranchStmt" -}}
{{ template "prefix1" }}BranchStmt
{{ template "prefix2" }}├── TokPos = {{ position .TokPos}}
{{ template "prefix2" }}├── Tok = {{ .Tok}}
{{ template "prefix2" }}└── Label
{{with .Label}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.CallExpr" -}}
{{ template "prefix1" }}CallExpr
{{ template "prefix2" }}├── Fun
{{with .Fun}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Lparen = {{ position .Lparen}}
{{ template "prefix2" }}├── Args = (length={{ len .Args }})
{{ range .Args }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Ellipsis = {{ position .Ellipsis}}
{{ template "prefix2" }}└── Rparen = {{ position .Rparen}}
{{ end }}
{{ define "*ast.CaseClause" -}}
{{ template "prefix1" }}CaseClause
{{ template "prefix2" }}├── Case = {{ position .Case}}
{{ template "prefix2" }}├── List = (length={{ len .List }})
{{ range .List }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Colon = {{ position .Colon}}
{{ template "prefix2" }}└── Body = (length={{ len .Body }})
{{ range .Body }}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.ChanType" -}}
{{ template "prefix1" }}ChanType
{{ template "prefix2" }}├── Begin = {{ position .Begin}}
{{ template "prefix2" }}├── Arrow = {{ position .Arrow}}
{{ template "prefix2" }}├── Dir = {{ .Dir}}
{{ template "prefix2" }}└── Value
{{with .Value}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.CommClause" -}}
{{ template "prefix1" }}CommClause
{{ template "prefix2" }}├── Case = {{ position .Case}}
{{ template "prefix2" }}├── Comm
{{with .Comm}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Colon = {{ position .Colon}}
{{ template "prefix2" }}└── Body = (length={{ len .Body }})
{{ range .Body }}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.Comment" -}}
{{ template "prefix1" }}Comment
{{ template "prefix2" }}├── Slash = {{ position .Slash}}
{{ template "prefix2" }}└── Text = {{ .Text}}
{{ end }}
{{ define "*ast.CommentGroup" -}}
{{ template "prefix1" }}CommentGroup
{{ template "prefix2" }}└── List = (length={{ len .List }})
{{ range .List }}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.CompositeLit" -}}
{{ template "prefix1" }}CompositeLit
{{ template "prefix2" }}├── Type
{{with .Type}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Lbrace = {{ position .Lbrace}}
{{ template "prefix2" }}├── Elts = (length={{ len .Elts }})
{{ range .Elts }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Rbrace = {{ position .Rbrace}}
{{ template "prefix2" }}└── Incomplete = {{ .Incomplete}}
{{ end }}
{{ define "*ast.DeclStmt" -}}
{{ template "prefix1" }}DeclStmt
{{ template "prefix2" }}└── Decl
{{with .Decl}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.DeferStmt" -}}
{{ template "prefix1" }}DeferStmt
{{ template "prefix2" }}├── Defer = {{ position .Defer}}
{{ template "prefix2" }}└── Call
{{with .Call}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.Ellipsis" -}}
{{ template "prefix1" }}Ellipsis
{{ template "prefix2" }}├── Ellipsis = {{ position .Ellipsis}}
{{ template "prefix2" }}└── Elt
{{with .Elt}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.EmptyStmt" -}}
{{ template "prefix1" }}EmptyStmt
{{ template "prefix2" }}├── Semicolon = {{ position .Semicolon}}
{{ template "prefix2" }}└── Implicit = {{ .Implicit}}
{{ end }}
{{ define "*ast.ExprStmt" -}}
{{ template "prefix1" }}ExprStmt
{{ template "prefix2" }}└── X
{{with .X}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.Field" -}}
{{ template "prefix1" }}Field
{{ template "prefix2" }}├── Doc
{{with .Doc}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Names = (length={{ len .Names }})
{{ range .Names }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Type
{{with .Type}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Tag
{{with .Tag}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Comment
{{with .Comment}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.FieldList" -}}
{{ template "prefix1" }}FieldList
{{ template "prefix2" }}├── Opening = {{ position .Opening}}
{{ template "prefix2" }}├── List = (length={{ len .List }})
{{ range .List }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Closing = {{ position .Closing}}
{{ end }}
{{ define "*ast.File" -}}
{{ template "prefix1" }}File
{{ template "prefix2" }}├── Doc
{{with .Doc}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Package = {{ position .Package}}
{{ template "prefix2" }}├── Name
{{with .Name}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Decls = (length={{ len .Decls }})
{{ range .Decls }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Scope
{{ template "prefix2" }}├── Imports = (length={{ len .Imports }})
{{ range .Imports }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Unresolved = (length={{ len .Unresolved }})
{{ range .Unresolved }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Comments = (length={{ len .Comments }})
{{ range .Comments }}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.ForStmt" -}}
{{ template "prefix1" }}ForStmt
{{ template "prefix2" }}├── For = {{ position .For}}
{{ template "prefix2" }}├── Init
{{with .Init}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Cond
{{with .Cond}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Post
{{with .Post}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Body
{{with .Body}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.FuncDecl" -}}
{{ template "prefix1" }}FuncDecl
{{ template "prefix2" }}├── Doc
{{with .Doc}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Recv
{{with .Recv}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Name
{{with .Name}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Type
{{with .Type}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Body
{{with .Body}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.FuncLit" -}}
{{ template "prefix1" }}FuncLit
{{ template "prefix2" }}├── Type
{{with .Type}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Body
{{with .Body}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.FuncType" -}}
{{ template "prefix1" }}FuncType
{{ template "prefix2" }}├── Func = {{ position .Func}}
{{ template "prefix2" }}├── Params
{{with .Params}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Results
{{with .Results}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.GenDecl" -}}
{{ template "prefix1" }}GenDecl
{{ template "prefix2" }}├── Doc
{{with .Doc}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── TokPos = {{ position .TokPos}}
{{ template "prefix2" }}├── Tok = {{ .Tok}}
{{ template "prefix2" }}├── Lparen = {{ position .Lparen}}
{{ template "prefix2" }}├── Specs = (length={{ len .Specs }})
{{ range .Specs }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Rparen = {{ position .Rparen}}
{{ end }}
{{ define "*ast.GoStmt" -}}
{{ template "prefix1" }}GoStmt
{{ template "prefix2" }}├── Go = {{ position .Go}}
{{ template "prefix2" }}└── Call
{{with .Call}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.Ident" -}}
{{ template "prefix1" }}Ident
{{ template "prefix2" }}├── NamePos = {{ position .NamePos}}
{{ template "prefix2" }}├── Name = {{ .Name}}
{{ template "prefix2" }}└── Obj
{{with .Obj}}{{ obj . }}{{ end -}}
{{ end }}
{{ define "*ast.IfStmt" -}}
{{ template "prefix1" }}IfStmt
{{ template "prefix2" }}├── If = {{ position .If}}
{{ template "prefix2" }}├── Init
{{with .Init}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Cond
{{with .Cond}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Body
{{with .Body}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Else
{{with .Else}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.ImportSpec" -}}
{{ template "prefix1" }}ImportSpec
{{ template "prefix2" }}├── Doc
{{with .Doc}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Name
{{with .Name}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Path
{{with .Path}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Comment
{{with .Comment}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── EndPos = {{ position .EndPos}}
{{ end }}
{{ define "*ast.IncDecStmt" -}}
{{ template "prefix1" }}IncDecStmt
{{ template "prefix2" }}├── X
{{with .X}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── TokPos = {{ position .TokPos}}
{{ template "prefix2" }}└── Tok = {{ .Tok}}
{{ end }}
{{ define "*ast.IndexExpr" -}}
{{ template "prefix1" }}IndexExpr
{{ template "prefix2" }}├── X
{{with .X}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Lbrack = {{ position .Lbrack}}
{{ template "prefix2" }}├── Index
{{with .Index}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Rbrack = {{ position .Rbrack}}
{{ end }}
{{ define "*ast.InterfaceType" -}}
{{ template "prefix1" }}InterfaceType
{{ template "prefix2" }}├── Interface = {{ position .Interface}}
{{ template "prefix2" }}├── Methods
{{with .Methods}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Incomplete = {{ .Incomplete}}
{{ end }}
{{ define "*ast.KeyValueExpr" -}}
{{ template "prefix1" }}KeyValueExpr
{{ template "prefix2" }}├── Key
{{with .Key}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Colon = {{ position .Colon}}
{{ template "prefix2" }}└── Value
{{with .Value}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.LabeledStmt" -}}
{{ template "prefix1" }}LabeledStmt
{{ template "prefix2" }}├── Label
{{with .Label}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Colon = {{ position .Colon}}
{{ template "prefix2" }}└── Stmt
{{with .Stmt}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.MapType" -}}
{{ template "prefix1" }}MapType
{{ template "prefix2" }}├── Map = {{ position .Map}}
{{ template "prefix2" }}├── Key
{{with .Key}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Value
{{with .Value}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.Package" -}}
{{ template "prefix1" }}Package
{{ template "prefix2" }}├── Name = {{ .Name}}
{{ template "prefix2" }}├── Scope
{{with .Scope}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.ParenExpr" -}}
{{ template "prefix1" }}ParenExpr
{{ template "prefix2" }}├── Lparen = {{ position .Lparen}}
{{ template "prefix2" }}├── X
{{with .X}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Rparen = {{ position .Rparen}}
{{ end }}
{{ define "*ast.RangeStmt" -}}
{{ template "prefix1" }}RangeStmt
{{ template "prefix2" }}├── For = {{ position .For}}
{{ template "prefix2" }}├── Key
{{with .Key}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Value
{{with .Value}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── TokPos = {{ position .TokPos}}
{{ template "prefix2" }}├── Tok = {{ .Tok}}
{{ template "prefix2" }}├── X
{{with .X}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Body
{{with .Body}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.ReturnStmt" -}}
{{ template "prefix1" }}ReturnStmt
{{ template "prefix2" }}├── Return = {{ position .Return}}
{{ template "prefix2" }}└── Results = (length={{ len .Results }})
{{ range .Results }}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.SelectStmt" -}}
{{ template "prefix1" }}SelectStmt
{{ template "prefix2" }}├── Select = {{ position .Select}}
{{ template "prefix2" }}└── Body
{{with .Body}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.SelectorExpr" -}}
{{ template "prefix1" }}SelectorExpr
{{ template "prefix2" }}├── X
{{with .X}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Sel
{{with .Sel}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.SendStmt" -}}
{{ template "prefix1" }}SendStmt
{{ template "prefix2" }}├── Chan
{{with .Chan}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Arrow = {{ position .Arrow}}
{{ template "prefix2" }}└── Value
{{with .Value}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.SliceExpr" -}}
{{ template "prefix1" }}SliceExpr
{{ template "prefix2" }}├── X
{{with .X}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Lbrack = {{ position .Lbrack}}
{{ template "prefix2" }}├── Low
{{with .Low}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── High
{{with .High}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Max
{{with .Max}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Slice3 = {{ .Slice3}}
{{ template "prefix2" }}└── Rbrack = {{ position .Rbrack}}
{{ end }}
{{ define "*ast.StarExpr" -}}
{{ template "prefix1" }}StarExpr
{{ template "prefix2" }}├── Star = {{ position .Star}}
{{ template "prefix2" }}└── X
{{with .X}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.StructType" -}}
{{ template "prefix1" }}StructType
{{ template "prefix2" }}├── Struct = {{ position .Struct}}
{{ template "prefix2" }}├── Fields
{{with .Fields}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Incomplete = {{ .Incomplete}}
{{ end }}
{{ define "*ast.SwitchStmt" -}}
{{ template "prefix1" }}SwitchStmt
{{ template "prefix2" }}├── Switch = {{ position .Switch}}
{{ template "prefix2" }}├── Init
{{with .Init}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Tag
{{with .Tag}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Body
{{with .Body}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.TypeAssertExpr" -}}
{{ template "prefix1" }}TypeAssertExpr
{{ template "prefix2" }}├── X
{{with .X}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Lparen = {{ position .Lparen}}
{{ template "prefix2" }}├── Type
{{with .Type}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Rparen = {{ position .Rparen}}
{{ end }}
{{ define "*ast.TypeSpec" -}}
{{ template "prefix1" }}TypeSpec
{{ template "prefix2" }}├── Doc
{{with .Doc}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Name
{{with .Name}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Assign = {{ position .Assign}}
{{ template "prefix2" }}├── Type
{{with .Type}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Comment
{{with .Comment}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.TypeSwitchStmt" -}}
{{ template "prefix1" }}TypeSwitchStmt
{{ template "prefix2" }}├── Switch = {{ position .Switch}}
{{ template "prefix2" }}├── Init
{{with .Init}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Assign
{{with .Assign}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Body
{{with .Body}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.UnaryExpr" -}}
{{ template "prefix1" }}UnaryExpr
{{ template "prefix2" }}├── OpPos = {{ position .OpPos}}
{{ template "prefix2" }}├── Op = {{ .Op}}
{{ template "prefix2" }}└── X
{{with .X}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.ValueSpec" -}}
{{ template "prefix1" }}ValueSpec
{{ template "prefix2" }}├── Doc
{{with .Doc}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Names = (length={{ len .Names }})
{{ range .Names }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Type
{{with .Type}}{{ tree . }}{{ end -}}
{{ template "prefix2" }}├── Values = (length={{ len .Values }})
{{ range .Values }}{{ tree . }}{{ end -}}
{{ template "prefix2" }}└── Comment
{{with .Comment}}{{ tree . }}{{ end -}}
{{ end }}
{{ define "*ast.Object" -}}
{{ template "prefix1" }}Object
{{ template "prefix2" }}├── Kind = {{ .Kind }}
{{ template "prefix2" }}├── Name = {{ .Name }}
{{ template "prefix2" }}├── Decl = {{ print .Decl }}
{{ template "prefix2" }}├── Data = {{ print .Data }}
{{ template "prefix2" }}└── Type = {{ print .Type }}
{{ end }}
