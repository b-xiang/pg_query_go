// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WindowDef) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("WindowDef")

	if node.EndOffset != nil {
		node.EndOffset.Fingerprint(ctx, "EndOffset")
	}

	ctx.WriteString(strconv.Itoa(int(node.FrameOptions)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	for _, subNode := range node.OrderClause {
		subNode.Fingerprint(ctx, "OrderClause")
	}

	for _, subNode := range node.PartitionClause {
		subNode.Fingerprint(ctx, "PartitionClause")
	}

	if node.Refname != nil {
		ctx.WriteString(*node.Refname)
	}

	if node.StartOffset != nil {
		node.StartOffset.Fingerprint(ctx, "StartOffset")
	}
}
