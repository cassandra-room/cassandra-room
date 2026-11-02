// Code generated from /Users/jungwook/cassandra-room/internal/cql/parser/Cql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Cql
import "github.com/antlr4-go/antlr/v4"

type BaseCqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCqlVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCqlStatements(ctx *CqlStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCqlStatement(ctx *CqlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUseStatement(ctx *UseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectors(ctx *SelectorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUnaliasedSelector(ctx *UnaliasedSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionAddition(ctx *SelectionAdditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionMultiplication(ctx *SelectionMultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionGroup(ctx *SelectionGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionGroupWithField(ctx *SelectionGroupWithFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectorModifier(ctx *SelectorModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitFieldSelectorModifier(ctx *FieldSelectorModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCollectionSubSelection(ctx *CollectionSubSelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionGroupWithoutField(ctx *SelectionGroupWithoutFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionTypeHint(ctx *SelectionTypeHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionList(ctx *SelectionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionMapOrSet(ctx *SelectionMapOrSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionMap(ctx *SelectionMapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionSet(ctx *SelectionSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionTupleOrNestedSelector(ctx *SelectionTupleOrNestedSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSimpleUnaliasedSelector(ctx *SimpleUnaliasedSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionFunction(ctx *SelectionFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionLiteral(ctx *SelectionLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSelectionFunctionArgs(ctx *SelectionFunctionArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSident(ctx *SidentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitRelationOrExpression(ctx *RelationOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCustomIndexExpression(ctx *CustomIndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitNormalInsertStatement(ctx *NormalInsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitJsonInsertStatement(ctx *JsonInsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitJsonValue(ctx *JsonValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUsingClause(ctx *UsingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUsingClauseObjective(ctx *UsingClauseObjectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUpdateConditions(ctx *UpdateConditionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDeleteSelection(ctx *DeleteSelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDeleteOp(ctx *DeleteOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUsingClauseDelete(ctx *UsingClauseDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitBatchStatement(ctx *BatchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitBatchStatementObjective(ctx *BatchStatementObjectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateAggregateStatement(ctx *CreateAggregateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropAggregateStatement(ctx *DropAggregateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateKeyspaceStatement(ctx *CreateKeyspaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTableDefinition(ctx *TableDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTableColumns(ctx *TableColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitColumnMask(ctx *ColumnMaskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitColumnMaskArguments(ctx *ColumnMaskArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTablePartitionKey(ctx *TablePartitionKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTableProperty(ctx *TablePropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTableClusteringOrder(ctx *TableClusteringOrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateTypeStatement(ctx *CreateTypeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTypeColumns(ctx *TypeColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitIndexIdent(ctx *IndexIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitViewPrimaryKey(ctx *ViewPrimaryKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitViewPartitionKey(ctx *ViewPartitionKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitViewProperty(ctx *ViewPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitViewClusteringOrder(ctx *ViewClusteringOrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateTriggerStatement(ctx *CreateTriggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropTriggerStatement(ctx *DropTriggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitAlterKeyspaceStatement(ctx *AlterKeyspaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitIsStaticColumn(ctx *IsStaticColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitAlterTypeStatement(ctx *AlterTypeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropKeyspaceStatement(ctx *DropKeyspaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropTableStatement(ctx *DropTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropTypeStatement(ctx *DropTypeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTruncateStatement(ctx *TruncateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitGrantPermissionsStatement(ctx *GrantPermissionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitRevokePermissionsStatement(ctx *RevokePermissionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitGrantRoleStatement(ctx *GrantRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitRevokeRoleStatement(ctx *RevokeRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitListPermissionsStatement(ctx *ListPermissionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitPermission(ctx *PermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitPermissionOrAll(ctx *PermissionOrAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitResource(ctx *ResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDataResource(ctx *DataResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitJmxResource(ctx *JmxResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitRoleResource(ctx *RoleResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitFunctionResource(ctx *FunctionResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateUserStatement(ctx *CreateUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitAlterUserStatement(ctx *AlterUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropUserStatement(ctx *DropUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitAddIdentityStatement(ctx *AddIdentityStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropIdentityStatement(ctx *DropIdentityStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitListUsersStatement(ctx *ListUsersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitAlterRoleStatement(ctx *AlterRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitListRolesStatement(ctx *ListRolesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitRoleOptions(ctx *RoleOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitRoleOption(ctx *RoleOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDcPermission(ctx *DcPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCidrPermission(ctx *CidrPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUserPassword(ctx *UserPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitDescribeStatement(ctx *DescribeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCident(ctx *CidentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitFident(ctx *FidentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitNoncol_ident(ctx *Noncol_identContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitKeyspaceName(ctx *KeyspaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitColumnFamilyName(ctx *ColumnFamilyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUserTypeName(ctx *UserTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUserOrRoleName(ctx *UserOrRoleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitKsName(ctx *KsNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCfName(ctx *CfNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitIdxName(ctx *IdxNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitRoleName(ctx *RoleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitFullMapLiteral(ctx *FullMapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSetOrMapLiteral(ctx *SetOrMapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSetLiteral(ctx *SetLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCollectionLiteral(ctx *CollectionLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitListLiteral(ctx *ListLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUsertypeLiteral(ctx *UsertypeLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTupleLiteral(ctx *TupleLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitIntValue(ctx *IntValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitAllowedFunctionName(ctx *AllowedFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTermAddition(ctx *TermAdditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTermMultiplication(ctx *TermMultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTermGroup(ctx *TermGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSimpleTerm(ctx *SimpleTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitColumnOperation(ctx *ColumnOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitColumnOperationDifferentiator(ctx *ColumnOperationDifferentiatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitNormalColumnOperation(ctx *NormalColumnOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitShorthandColumnOperation(ctx *ShorthandColumnOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCollectionColumnOperation(ctx *CollectionColumnOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUdtColumnOperation(ctx *UdtColumnOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitColumnCondition(ctx *ColumnConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitProperties(ctx *PropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitPropertyValue(ctx *PropertyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitRelationType(ctx *RelationTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitContainsOperator(ctx *ContainsOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitInMarker(ctx *InMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTupleOfIdentifiers(ctx *TupleOfIdentifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitSingleColumnInValues(ctx *SingleColumnInValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTupleOfTupleLiterals(ctx *TupleOfTupleLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitMarkerForTuple(ctx *MarkerForTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTupleOfMarkersForTuples(ctx *TupleOfMarkersForTuplesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitInMarkerForTuple(ctx *InMarkerForTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitComparatorType(ctx *ComparatorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitNative_type(ctx *Native_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitCollection_type(ctx *Collection_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitTuple_type(ctx *Tuple_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitVector_type(ctx *Vector_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUsername(ctx *UsernameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitIdentity(ctx *IdentityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitMbean(ctx *MbeanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitNon_type_ident(ctx *Non_type_identContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUnreserved_keyword(ctx *Unreserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitUnreserved_function_keyword(ctx *Unreserved_function_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCqlVisitor) VisitBasic_unreserved_keyword(ctx *Basic_unreserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}
