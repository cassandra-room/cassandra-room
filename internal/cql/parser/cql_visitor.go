// Code generated from /Users/jungwook/cassandra-room/internal/cql/parser/Cql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Cql
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CqlParser.
type CqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CqlParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by CqlParser#cqlStatements.
	VisitCqlStatements(ctx *CqlStatementsContext) interface{}

	// Visit a parse tree produced by CqlParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by CqlParser#cqlStatement.
	VisitCqlStatement(ctx *CqlStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#useStatement.
	VisitUseStatement(ctx *UseStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by CqlParser#selectors.
	VisitSelectors(ctx *SelectorsContext) interface{}

	// Visit a parse tree produced by CqlParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by CqlParser#unaliasedSelector.
	VisitUnaliasedSelector(ctx *UnaliasedSelectorContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionAddition.
	VisitSelectionAddition(ctx *SelectionAdditionContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionMultiplication.
	VisitSelectionMultiplication(ctx *SelectionMultiplicationContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionGroup.
	VisitSelectionGroup(ctx *SelectionGroupContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionGroupWithField.
	VisitSelectionGroupWithField(ctx *SelectionGroupWithFieldContext) interface{}

	// Visit a parse tree produced by CqlParser#selectorModifier.
	VisitSelectorModifier(ctx *SelectorModifierContext) interface{}

	// Visit a parse tree produced by CqlParser#fieldSelectorModifier.
	VisitFieldSelectorModifier(ctx *FieldSelectorModifierContext) interface{}

	// Visit a parse tree produced by CqlParser#collectionSubSelection.
	VisitCollectionSubSelection(ctx *CollectionSubSelectionContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionGroupWithoutField.
	VisitSelectionGroupWithoutField(ctx *SelectionGroupWithoutFieldContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionTypeHint.
	VisitSelectionTypeHint(ctx *SelectionTypeHintContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionList.
	VisitSelectionList(ctx *SelectionListContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionMapOrSet.
	VisitSelectionMapOrSet(ctx *SelectionMapOrSetContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionMap.
	VisitSelectionMap(ctx *SelectionMapContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionSet.
	VisitSelectionSet(ctx *SelectionSetContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionTupleOrNestedSelector.
	VisitSelectionTupleOrNestedSelector(ctx *SelectionTupleOrNestedSelectorContext) interface{}

	// Visit a parse tree produced by CqlParser#simpleUnaliasedSelector.
	VisitSimpleUnaliasedSelector(ctx *SimpleUnaliasedSelectorContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionFunction.
	VisitSelectionFunction(ctx *SelectionFunctionContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionLiteral.
	VisitSelectionLiteral(ctx *SelectionLiteralContext) interface{}

	// Visit a parse tree produced by CqlParser#selectionFunctionArgs.
	VisitSelectionFunctionArgs(ctx *SelectionFunctionArgsContext) interface{}

	// Visit a parse tree produced by CqlParser#sident.
	VisitSident(ctx *SidentContext) interface{}

	// Visit a parse tree produced by CqlParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by CqlParser#relationOrExpression.
	VisitRelationOrExpression(ctx *RelationOrExpressionContext) interface{}

	// Visit a parse tree produced by CqlParser#customIndexExpression.
	VisitCustomIndexExpression(ctx *CustomIndexExpressionContext) interface{}

	// Visit a parse tree produced by CqlParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by CqlParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by CqlParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#normalInsertStatement.
	VisitNormalInsertStatement(ctx *NormalInsertStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#jsonInsertStatement.
	VisitJsonInsertStatement(ctx *JsonInsertStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#jsonValue.
	VisitJsonValue(ctx *JsonValueContext) interface{}

	// Visit a parse tree produced by CqlParser#usingClause.
	VisitUsingClause(ctx *UsingClauseContext) interface{}

	// Visit a parse tree produced by CqlParser#usingClauseObjective.
	VisitUsingClauseObjective(ctx *UsingClauseObjectiveContext) interface{}

	// Visit a parse tree produced by CqlParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#updateConditions.
	VisitUpdateConditions(ctx *UpdateConditionsContext) interface{}

	// Visit a parse tree produced by CqlParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#deleteSelection.
	VisitDeleteSelection(ctx *DeleteSelectionContext) interface{}

	// Visit a parse tree produced by CqlParser#deleteOp.
	VisitDeleteOp(ctx *DeleteOpContext) interface{}

	// Visit a parse tree produced by CqlParser#usingClauseDelete.
	VisitUsingClauseDelete(ctx *UsingClauseDeleteContext) interface{}

	// Visit a parse tree produced by CqlParser#batchStatement.
	VisitBatchStatement(ctx *BatchStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#batchStatementObjective.
	VisitBatchStatementObjective(ctx *BatchStatementObjectiveContext) interface{}

	// Visit a parse tree produced by CqlParser#createAggregateStatement.
	VisitCreateAggregateStatement(ctx *CreateAggregateStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropAggregateStatement.
	VisitDropAggregateStatement(ctx *DropAggregateStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#createFunctionStatement.
	VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropFunctionStatement.
	VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#createKeyspaceStatement.
	VisitCreateKeyspaceStatement(ctx *CreateKeyspaceStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#createTableStatement.
	VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#tableDefinition.
	VisitTableDefinition(ctx *TableDefinitionContext) interface{}

	// Visit a parse tree produced by CqlParser#tableColumns.
	VisitTableColumns(ctx *TableColumnsContext) interface{}

	// Visit a parse tree produced by CqlParser#columnMask.
	VisitColumnMask(ctx *ColumnMaskContext) interface{}

	// Visit a parse tree produced by CqlParser#columnMaskArguments.
	VisitColumnMaskArguments(ctx *ColumnMaskArgumentsContext) interface{}

	// Visit a parse tree produced by CqlParser#tablePartitionKey.
	VisitTablePartitionKey(ctx *TablePartitionKeyContext) interface{}

	// Visit a parse tree produced by CqlParser#tableProperty.
	VisitTableProperty(ctx *TablePropertyContext) interface{}

	// Visit a parse tree produced by CqlParser#tableClusteringOrder.
	VisitTableClusteringOrder(ctx *TableClusteringOrderContext) interface{}

	// Visit a parse tree produced by CqlParser#createTypeStatement.
	VisitCreateTypeStatement(ctx *CreateTypeStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#typeColumns.
	VisitTypeColumns(ctx *TypeColumnsContext) interface{}

	// Visit a parse tree produced by CqlParser#createIndexStatement.
	VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#indexIdent.
	VisitIndexIdent(ctx *IndexIdentContext) interface{}

	// Visit a parse tree produced by CqlParser#createMaterializedViewStatement.
	VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#viewPrimaryKey.
	VisitViewPrimaryKey(ctx *ViewPrimaryKeyContext) interface{}

	// Visit a parse tree produced by CqlParser#viewPartitionKey.
	VisitViewPartitionKey(ctx *ViewPartitionKeyContext) interface{}

	// Visit a parse tree produced by CqlParser#viewProperty.
	VisitViewProperty(ctx *ViewPropertyContext) interface{}

	// Visit a parse tree produced by CqlParser#viewClusteringOrder.
	VisitViewClusteringOrder(ctx *ViewClusteringOrderContext) interface{}

	// Visit a parse tree produced by CqlParser#createTriggerStatement.
	VisitCreateTriggerStatement(ctx *CreateTriggerStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropTriggerStatement.
	VisitDropTriggerStatement(ctx *DropTriggerStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#alterKeyspaceStatement.
	VisitAlterKeyspaceStatement(ctx *AlterKeyspaceStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#alterTableStatement.
	VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#isStaticColumn.
	VisitIsStaticColumn(ctx *IsStaticColumnContext) interface{}

	// Visit a parse tree produced by CqlParser#alterMaterializedViewStatement.
	VisitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#alterTypeStatement.
	VisitAlterTypeStatement(ctx *AlterTypeStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropKeyspaceStatement.
	VisitDropKeyspaceStatement(ctx *DropKeyspaceStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropTableStatement.
	VisitDropTableStatement(ctx *DropTableStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropTypeStatement.
	VisitDropTypeStatement(ctx *DropTypeStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropIndexStatement.
	VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropMaterializedViewStatement.
	VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#truncateStatement.
	VisitTruncateStatement(ctx *TruncateStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#grantPermissionsStatement.
	VisitGrantPermissionsStatement(ctx *GrantPermissionsStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#revokePermissionsStatement.
	VisitRevokePermissionsStatement(ctx *RevokePermissionsStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#grantRoleStatement.
	VisitGrantRoleStatement(ctx *GrantRoleStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#revokeRoleStatement.
	VisitRevokeRoleStatement(ctx *RevokeRoleStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#listPermissionsStatement.
	VisitListPermissionsStatement(ctx *ListPermissionsStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#permission.
	VisitPermission(ctx *PermissionContext) interface{}

	// Visit a parse tree produced by CqlParser#permissionOrAll.
	VisitPermissionOrAll(ctx *PermissionOrAllContext) interface{}

	// Visit a parse tree produced by CqlParser#resource.
	VisitResource(ctx *ResourceContext) interface{}

	// Visit a parse tree produced by CqlParser#dataResource.
	VisitDataResource(ctx *DataResourceContext) interface{}

	// Visit a parse tree produced by CqlParser#jmxResource.
	VisitJmxResource(ctx *JmxResourceContext) interface{}

	// Visit a parse tree produced by CqlParser#roleResource.
	VisitRoleResource(ctx *RoleResourceContext) interface{}

	// Visit a parse tree produced by CqlParser#functionResource.
	VisitFunctionResource(ctx *FunctionResourceContext) interface{}

	// Visit a parse tree produced by CqlParser#createUserStatement.
	VisitCreateUserStatement(ctx *CreateUserStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#alterUserStatement.
	VisitAlterUserStatement(ctx *AlterUserStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropUserStatement.
	VisitDropUserStatement(ctx *DropUserStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#addIdentityStatement.
	VisitAddIdentityStatement(ctx *AddIdentityStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropIdentityStatement.
	VisitDropIdentityStatement(ctx *DropIdentityStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#listUsersStatement.
	VisitListUsersStatement(ctx *ListUsersStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#createRoleStatement.
	VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#alterRoleStatement.
	VisitAlterRoleStatement(ctx *AlterRoleStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#dropRoleStatement.
	VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#listRolesStatement.
	VisitListRolesStatement(ctx *ListRolesStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#roleOptions.
	VisitRoleOptions(ctx *RoleOptionsContext) interface{}

	// Visit a parse tree produced by CqlParser#roleOption.
	VisitRoleOption(ctx *RoleOptionContext) interface{}

	// Visit a parse tree produced by CqlParser#dcPermission.
	VisitDcPermission(ctx *DcPermissionContext) interface{}

	// Visit a parse tree produced by CqlParser#cidrPermission.
	VisitCidrPermission(ctx *CidrPermissionContext) interface{}

	// Visit a parse tree produced by CqlParser#userPassword.
	VisitUserPassword(ctx *UserPasswordContext) interface{}

	// Visit a parse tree produced by CqlParser#describeStatement.
	VisitDescribeStatement(ctx *DescribeStatementContext) interface{}

	// Visit a parse tree produced by CqlParser#cident.
	VisitCident(ctx *CidentContext) interface{}

	// Visit a parse tree produced by CqlParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by CqlParser#fident.
	VisitFident(ctx *FidentContext) interface{}

	// Visit a parse tree produced by CqlParser#noncol_ident.
	VisitNoncol_ident(ctx *Noncol_identContext) interface{}

	// Visit a parse tree produced by CqlParser#keyspaceName.
	VisitKeyspaceName(ctx *KeyspaceNameContext) interface{}

	// Visit a parse tree produced by CqlParser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by CqlParser#columnFamilyName.
	VisitColumnFamilyName(ctx *ColumnFamilyNameContext) interface{}

	// Visit a parse tree produced by CqlParser#userTypeName.
	VisitUserTypeName(ctx *UserTypeNameContext) interface{}

	// Visit a parse tree produced by CqlParser#userOrRoleName.
	VisitUserOrRoleName(ctx *UserOrRoleNameContext) interface{}

	// Visit a parse tree produced by CqlParser#ksName.
	VisitKsName(ctx *KsNameContext) interface{}

	// Visit a parse tree produced by CqlParser#cfName.
	VisitCfName(ctx *CfNameContext) interface{}

	// Visit a parse tree produced by CqlParser#idxName.
	VisitIdxName(ctx *IdxNameContext) interface{}

	// Visit a parse tree produced by CqlParser#roleName.
	VisitRoleName(ctx *RoleNameContext) interface{}

	// Visit a parse tree produced by CqlParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by CqlParser#fullMapLiteral.
	VisitFullMapLiteral(ctx *FullMapLiteralContext) interface{}

	// Visit a parse tree produced by CqlParser#setOrMapLiteral.
	VisitSetOrMapLiteral(ctx *SetOrMapLiteralContext) interface{}

	// Visit a parse tree produced by CqlParser#setLiteral.
	VisitSetLiteral(ctx *SetLiteralContext) interface{}

	// Visit a parse tree produced by CqlParser#mapLiteral.
	VisitMapLiteral(ctx *MapLiteralContext) interface{}

	// Visit a parse tree produced by CqlParser#collectionLiteral.
	VisitCollectionLiteral(ctx *CollectionLiteralContext) interface{}

	// Visit a parse tree produced by CqlParser#listLiteral.
	VisitListLiteral(ctx *ListLiteralContext) interface{}

	// Visit a parse tree produced by CqlParser#usertypeLiteral.
	VisitUsertypeLiteral(ctx *UsertypeLiteralContext) interface{}

	// Visit a parse tree produced by CqlParser#tupleLiteral.
	VisitTupleLiteral(ctx *TupleLiteralContext) interface{}

	// Visit a parse tree produced by CqlParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by CqlParser#intValue.
	VisitIntValue(ctx *IntValueContext) interface{}

	// Visit a parse tree produced by CqlParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by CqlParser#allowedFunctionName.
	VisitAllowedFunctionName(ctx *AllowedFunctionNameContext) interface{}

	// Visit a parse tree produced by CqlParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by CqlParser#functionArgs.
	VisitFunctionArgs(ctx *FunctionArgsContext) interface{}

	// Visit a parse tree produced by CqlParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by CqlParser#termAddition.
	VisitTermAddition(ctx *TermAdditionContext) interface{}

	// Visit a parse tree produced by CqlParser#termMultiplication.
	VisitTermMultiplication(ctx *TermMultiplicationContext) interface{}

	// Visit a parse tree produced by CqlParser#termGroup.
	VisitTermGroup(ctx *TermGroupContext) interface{}

	// Visit a parse tree produced by CqlParser#simpleTerm.
	VisitSimpleTerm(ctx *SimpleTermContext) interface{}

	// Visit a parse tree produced by CqlParser#columnOperation.
	VisitColumnOperation(ctx *ColumnOperationContext) interface{}

	// Visit a parse tree produced by CqlParser#columnOperationDifferentiator.
	VisitColumnOperationDifferentiator(ctx *ColumnOperationDifferentiatorContext) interface{}

	// Visit a parse tree produced by CqlParser#normalColumnOperation.
	VisitNormalColumnOperation(ctx *NormalColumnOperationContext) interface{}

	// Visit a parse tree produced by CqlParser#shorthandColumnOperation.
	VisitShorthandColumnOperation(ctx *ShorthandColumnOperationContext) interface{}

	// Visit a parse tree produced by CqlParser#collectionColumnOperation.
	VisitCollectionColumnOperation(ctx *CollectionColumnOperationContext) interface{}

	// Visit a parse tree produced by CqlParser#udtColumnOperation.
	VisitUdtColumnOperation(ctx *UdtColumnOperationContext) interface{}

	// Visit a parse tree produced by CqlParser#columnCondition.
	VisitColumnCondition(ctx *ColumnConditionContext) interface{}

	// Visit a parse tree produced by CqlParser#properties.
	VisitProperties(ctx *PropertiesContext) interface{}

	// Visit a parse tree produced by CqlParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by CqlParser#propertyValue.
	VisitPropertyValue(ctx *PropertyValueContext) interface{}

	// Visit a parse tree produced by CqlParser#relationType.
	VisitRelationType(ctx *RelationTypeContext) interface{}

	// Visit a parse tree produced by CqlParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by CqlParser#containsOperator.
	VisitContainsOperator(ctx *ContainsOperatorContext) interface{}

	// Visit a parse tree produced by CqlParser#inMarker.
	VisitInMarker(ctx *InMarkerContext) interface{}

	// Visit a parse tree produced by CqlParser#tupleOfIdentifiers.
	VisitTupleOfIdentifiers(ctx *TupleOfIdentifiersContext) interface{}

	// Visit a parse tree produced by CqlParser#singleColumnInValues.
	VisitSingleColumnInValues(ctx *SingleColumnInValuesContext) interface{}

	// Visit a parse tree produced by CqlParser#tupleOfTupleLiterals.
	VisitTupleOfTupleLiterals(ctx *TupleOfTupleLiteralsContext) interface{}

	// Visit a parse tree produced by CqlParser#markerForTuple.
	VisitMarkerForTuple(ctx *MarkerForTupleContext) interface{}

	// Visit a parse tree produced by CqlParser#tupleOfMarkersForTuples.
	VisitTupleOfMarkersForTuples(ctx *TupleOfMarkersForTuplesContext) interface{}

	// Visit a parse tree produced by CqlParser#inMarkerForTuple.
	VisitInMarkerForTuple(ctx *InMarkerForTupleContext) interface{}

	// Visit a parse tree produced by CqlParser#comparatorType.
	VisitComparatorType(ctx *ComparatorTypeContext) interface{}

	// Visit a parse tree produced by CqlParser#native_type.
	VisitNative_type(ctx *Native_typeContext) interface{}

	// Visit a parse tree produced by CqlParser#collection_type.
	VisitCollection_type(ctx *Collection_typeContext) interface{}

	// Visit a parse tree produced by CqlParser#tuple_type.
	VisitTuple_type(ctx *Tuple_typeContext) interface{}

	// Visit a parse tree produced by CqlParser#vector_type.
	VisitVector_type(ctx *Vector_typeContext) interface{}

	// Visit a parse tree produced by CqlParser#username.
	VisitUsername(ctx *UsernameContext) interface{}

	// Visit a parse tree produced by CqlParser#identity.
	VisitIdentity(ctx *IdentityContext) interface{}

	// Visit a parse tree produced by CqlParser#mbean.
	VisitMbean(ctx *MbeanContext) interface{}

	// Visit a parse tree produced by CqlParser#non_type_ident.
	VisitNon_type_ident(ctx *Non_type_identContext) interface{}

	// Visit a parse tree produced by CqlParser#unreserved_keyword.
	VisitUnreserved_keyword(ctx *Unreserved_keywordContext) interface{}

	// Visit a parse tree produced by CqlParser#unreserved_function_keyword.
	VisitUnreserved_function_keyword(ctx *Unreserved_function_keywordContext) interface{}

	// Visit a parse tree produced by CqlParser#basic_unreserved_keyword.
	VisitBasic_unreserved_keyword(ctx *Basic_unreserved_keywordContext) interface{}
}
