// Auto-generated - DO NOT EDIT

package pg_query

func (node A_Indices) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_Indices")

	if node.Lidx != nil {
		node.Lidx.Fingerprint(ctx, "Lidx")
	}

	if node.Uidx != nil {
		node.Uidx.Fingerprint(ctx, "Uidx")
	}
}
