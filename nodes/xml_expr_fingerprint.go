// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node XmlExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("XmlExpr")

	for _, subNode := range node.ArgNames {
		subNode.Fingerprint(ctx, "ArgNames")
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx, "Args")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	for _, subNode := range node.NamedArgs {
		subNode.Fingerprint(ctx, "NamedArgs")
	}

	ctx.WriteString(strconv.Itoa(int(node.Op)))
	ctx.WriteString(strconv.Itoa(int(node.Type)))
	ctx.WriteString(strconv.Itoa(int(node.Typmod)))
	ctx.WriteString(strconv.Itoa(int(node.Xmloption)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
