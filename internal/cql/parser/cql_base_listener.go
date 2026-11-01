// Code generated from /Users/jungwook/cassandra-room/internal/cql/parser/Cql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Cql
import "github.com/antlr4-go/antlr/v4"

// BaseCqlListener is a complete listener for a parse tree produced by CqlParser.
type BaseCqlListener struct{}

var _ CqlListener = &BaseCqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseCqlListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseCqlListener) ExitRoot(ctx *RootContext) {}

// EnterCqlStatements is called when production cqlStatements is entered.
func (s *BaseCqlListener) EnterCqlStatements(ctx *CqlStatementsContext) {}

// ExitCqlStatements is called when production cqlStatements is exited.
func (s *BaseCqlListener) ExitCqlStatements(ctx *CqlStatementsContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseCqlListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseCqlListener) ExitQuery(ctx *QueryContext) {}

// EnterCqlStatement is called when production cqlStatement is entered.
func (s *BaseCqlListener) EnterCqlStatement(ctx *CqlStatementContext) {}

// ExitCqlStatement is called when production cqlStatement is exited.
func (s *BaseCqlListener) ExitCqlStatement(ctx *CqlStatementContext) {}

// EnterUseStatement is called when production useStatement is entered.
func (s *BaseCqlListener) EnterUseStatement(ctx *UseStatementContext) {}

// ExitUseStatement is called when production useStatement is exited.
func (s *BaseCqlListener) ExitUseStatement(ctx *UseStatementContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseCqlListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseCqlListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseCqlListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseCqlListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterSelectors is called when production selectors is entered.
func (s *BaseCqlListener) EnterSelectors(ctx *SelectorsContext) {}

// ExitSelectors is called when production selectors is exited.
func (s *BaseCqlListener) ExitSelectors(ctx *SelectorsContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseCqlListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseCqlListener) ExitSelector(ctx *SelectorContext) {}

// EnterUnaliasedSelector is called when production unaliasedSelector is entered.
func (s *BaseCqlListener) EnterUnaliasedSelector(ctx *UnaliasedSelectorContext) {}

// ExitUnaliasedSelector is called when production unaliasedSelector is exited.
func (s *BaseCqlListener) ExitUnaliasedSelector(ctx *UnaliasedSelectorContext) {}

// EnterSelectionAddition is called when production selectionAddition is entered.
func (s *BaseCqlListener) EnterSelectionAddition(ctx *SelectionAdditionContext) {}

// ExitSelectionAddition is called when production selectionAddition is exited.
func (s *BaseCqlListener) ExitSelectionAddition(ctx *SelectionAdditionContext) {}

// EnterSelectionMultiplication is called when production selectionMultiplication is entered.
func (s *BaseCqlListener) EnterSelectionMultiplication(ctx *SelectionMultiplicationContext) {}

// ExitSelectionMultiplication is called when production selectionMultiplication is exited.
func (s *BaseCqlListener) ExitSelectionMultiplication(ctx *SelectionMultiplicationContext) {}

// EnterSelectionGroup is called when production selectionGroup is entered.
func (s *BaseCqlListener) EnterSelectionGroup(ctx *SelectionGroupContext) {}

// ExitSelectionGroup is called when production selectionGroup is exited.
func (s *BaseCqlListener) ExitSelectionGroup(ctx *SelectionGroupContext) {}

// EnterSelectionGroupWithField is called when production selectionGroupWithField is entered.
func (s *BaseCqlListener) EnterSelectionGroupWithField(ctx *SelectionGroupWithFieldContext) {}

// ExitSelectionGroupWithField is called when production selectionGroupWithField is exited.
func (s *BaseCqlListener) ExitSelectionGroupWithField(ctx *SelectionGroupWithFieldContext) {}

// EnterSelectorModifier is called when production selectorModifier is entered.
func (s *BaseCqlListener) EnterSelectorModifier(ctx *SelectorModifierContext) {}

// ExitSelectorModifier is called when production selectorModifier is exited.
func (s *BaseCqlListener) ExitSelectorModifier(ctx *SelectorModifierContext) {}

// EnterFieldSelectorModifier is called when production fieldSelectorModifier is entered.
func (s *BaseCqlListener) EnterFieldSelectorModifier(ctx *FieldSelectorModifierContext) {}

// ExitFieldSelectorModifier is called when production fieldSelectorModifier is exited.
func (s *BaseCqlListener) ExitFieldSelectorModifier(ctx *FieldSelectorModifierContext) {}

// EnterCollectionSubSelection is called when production collectionSubSelection is entered.
func (s *BaseCqlListener) EnterCollectionSubSelection(ctx *CollectionSubSelectionContext) {}

// ExitCollectionSubSelection is called when production collectionSubSelection is exited.
func (s *BaseCqlListener) ExitCollectionSubSelection(ctx *CollectionSubSelectionContext) {}

// EnterSelectionGroupWithoutField is called when production selectionGroupWithoutField is entered.
func (s *BaseCqlListener) EnterSelectionGroupWithoutField(ctx *SelectionGroupWithoutFieldContext) {}

// ExitSelectionGroupWithoutField is called when production selectionGroupWithoutField is exited.
func (s *BaseCqlListener) ExitSelectionGroupWithoutField(ctx *SelectionGroupWithoutFieldContext) {}

// EnterSelectionTypeHint is called when production selectionTypeHint is entered.
func (s *BaseCqlListener) EnterSelectionTypeHint(ctx *SelectionTypeHintContext) {}

// ExitSelectionTypeHint is called when production selectionTypeHint is exited.
func (s *BaseCqlListener) ExitSelectionTypeHint(ctx *SelectionTypeHintContext) {}

// EnterSelectionList is called when production selectionList is entered.
func (s *BaseCqlListener) EnterSelectionList(ctx *SelectionListContext) {}

// ExitSelectionList is called when production selectionList is exited.
func (s *BaseCqlListener) ExitSelectionList(ctx *SelectionListContext) {}

// EnterSelectionMapOrSet is called when production selectionMapOrSet is entered.
func (s *BaseCqlListener) EnterSelectionMapOrSet(ctx *SelectionMapOrSetContext) {}

// ExitSelectionMapOrSet is called when production selectionMapOrSet is exited.
func (s *BaseCqlListener) ExitSelectionMapOrSet(ctx *SelectionMapOrSetContext) {}

// EnterSelectionMap is called when production selectionMap is entered.
func (s *BaseCqlListener) EnterSelectionMap(ctx *SelectionMapContext) {}

// ExitSelectionMap is called when production selectionMap is exited.
func (s *BaseCqlListener) ExitSelectionMap(ctx *SelectionMapContext) {}

// EnterSelectionSet is called when production selectionSet is entered.
func (s *BaseCqlListener) EnterSelectionSet(ctx *SelectionSetContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (s *BaseCqlListener) ExitSelectionSet(ctx *SelectionSetContext) {}

// EnterSelectionTupleOrNestedSelector is called when production selectionTupleOrNestedSelector is entered.
func (s *BaseCqlListener) EnterSelectionTupleOrNestedSelector(ctx *SelectionTupleOrNestedSelectorContext) {
}

// ExitSelectionTupleOrNestedSelector is called when production selectionTupleOrNestedSelector is exited.
func (s *BaseCqlListener) ExitSelectionTupleOrNestedSelector(ctx *SelectionTupleOrNestedSelectorContext) {
}

// EnterSimpleUnaliasedSelector is called when production simpleUnaliasedSelector is entered.
func (s *BaseCqlListener) EnterSimpleUnaliasedSelector(ctx *SimpleUnaliasedSelectorContext) {}

// ExitSimpleUnaliasedSelector is called when production simpleUnaliasedSelector is exited.
func (s *BaseCqlListener) ExitSimpleUnaliasedSelector(ctx *SimpleUnaliasedSelectorContext) {}

// EnterSelectionFunction is called when production selectionFunction is entered.
func (s *BaseCqlListener) EnterSelectionFunction(ctx *SelectionFunctionContext) {}

// ExitSelectionFunction is called when production selectionFunction is exited.
func (s *BaseCqlListener) ExitSelectionFunction(ctx *SelectionFunctionContext) {}

// EnterSelectionLiteral is called when production selectionLiteral is entered.
func (s *BaseCqlListener) EnterSelectionLiteral(ctx *SelectionLiteralContext) {}

// ExitSelectionLiteral is called when production selectionLiteral is exited.
func (s *BaseCqlListener) ExitSelectionLiteral(ctx *SelectionLiteralContext) {}

// EnterSelectionFunctionArgs is called when production selectionFunctionArgs is entered.
func (s *BaseCqlListener) EnterSelectionFunctionArgs(ctx *SelectionFunctionArgsContext) {}

// ExitSelectionFunctionArgs is called when production selectionFunctionArgs is exited.
func (s *BaseCqlListener) ExitSelectionFunctionArgs(ctx *SelectionFunctionArgsContext) {}

// EnterSident is called when production sident is entered.
func (s *BaseCqlListener) EnterSident(ctx *SidentContext) {}

// ExitSident is called when production sident is exited.
func (s *BaseCqlListener) ExitSident(ctx *SidentContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseCqlListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseCqlListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterRelationOrExpression is called when production relationOrExpression is entered.
func (s *BaseCqlListener) EnterRelationOrExpression(ctx *RelationOrExpressionContext) {}

// ExitRelationOrExpression is called when production relationOrExpression is exited.
func (s *BaseCqlListener) ExitRelationOrExpression(ctx *RelationOrExpressionContext) {}

// EnterCustomIndexExpression is called when production customIndexExpression is entered.
func (s *BaseCqlListener) EnterCustomIndexExpression(ctx *CustomIndexExpressionContext) {}

// ExitCustomIndexExpression is called when production customIndexExpression is exited.
func (s *BaseCqlListener) ExitCustomIndexExpression(ctx *CustomIndexExpressionContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseCqlListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseCqlListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseCqlListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseCqlListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseCqlListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseCqlListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterNormalInsertStatement is called when production normalInsertStatement is entered.
func (s *BaseCqlListener) EnterNormalInsertStatement(ctx *NormalInsertStatementContext) {}

// ExitNormalInsertStatement is called when production normalInsertStatement is exited.
func (s *BaseCqlListener) ExitNormalInsertStatement(ctx *NormalInsertStatementContext) {}

// EnterJsonInsertStatement is called when production jsonInsertStatement is entered.
func (s *BaseCqlListener) EnterJsonInsertStatement(ctx *JsonInsertStatementContext) {}

// ExitJsonInsertStatement is called when production jsonInsertStatement is exited.
func (s *BaseCqlListener) ExitJsonInsertStatement(ctx *JsonInsertStatementContext) {}

// EnterJsonValue is called when production jsonValue is entered.
func (s *BaseCqlListener) EnterJsonValue(ctx *JsonValueContext) {}

// ExitJsonValue is called when production jsonValue is exited.
func (s *BaseCqlListener) ExitJsonValue(ctx *JsonValueContext) {}

// EnterUsingClause is called when production usingClause is entered.
func (s *BaseCqlListener) EnterUsingClause(ctx *UsingClauseContext) {}

// ExitUsingClause is called when production usingClause is exited.
func (s *BaseCqlListener) ExitUsingClause(ctx *UsingClauseContext) {}

// EnterUsingClauseObjective is called when production usingClauseObjective is entered.
func (s *BaseCqlListener) EnterUsingClauseObjective(ctx *UsingClauseObjectiveContext) {}

// ExitUsingClauseObjective is called when production usingClauseObjective is exited.
func (s *BaseCqlListener) ExitUsingClauseObjective(ctx *UsingClauseObjectiveContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseCqlListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseCqlListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterUpdateConditions is called when production updateConditions is entered.
func (s *BaseCqlListener) EnterUpdateConditions(ctx *UpdateConditionsContext) {}

// ExitUpdateConditions is called when production updateConditions is exited.
func (s *BaseCqlListener) ExitUpdateConditions(ctx *UpdateConditionsContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseCqlListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseCqlListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterDeleteSelection is called when production deleteSelection is entered.
func (s *BaseCqlListener) EnterDeleteSelection(ctx *DeleteSelectionContext) {}

// ExitDeleteSelection is called when production deleteSelection is exited.
func (s *BaseCqlListener) ExitDeleteSelection(ctx *DeleteSelectionContext) {}

// EnterDeleteOp is called when production deleteOp is entered.
func (s *BaseCqlListener) EnterDeleteOp(ctx *DeleteOpContext) {}

// ExitDeleteOp is called when production deleteOp is exited.
func (s *BaseCqlListener) ExitDeleteOp(ctx *DeleteOpContext) {}

// EnterUsingClauseDelete is called when production usingClauseDelete is entered.
func (s *BaseCqlListener) EnterUsingClauseDelete(ctx *UsingClauseDeleteContext) {}

// ExitUsingClauseDelete is called when production usingClauseDelete is exited.
func (s *BaseCqlListener) ExitUsingClauseDelete(ctx *UsingClauseDeleteContext) {}

// EnterBatchStatement is called when production batchStatement is entered.
func (s *BaseCqlListener) EnterBatchStatement(ctx *BatchStatementContext) {}

// ExitBatchStatement is called when production batchStatement is exited.
func (s *BaseCqlListener) ExitBatchStatement(ctx *BatchStatementContext) {}

// EnterBatchStatementObjective is called when production batchStatementObjective is entered.
func (s *BaseCqlListener) EnterBatchStatementObjective(ctx *BatchStatementObjectiveContext) {}

// ExitBatchStatementObjective is called when production batchStatementObjective is exited.
func (s *BaseCqlListener) ExitBatchStatementObjective(ctx *BatchStatementObjectiveContext) {}

// EnterCreateAggregateStatement is called when production createAggregateStatement is entered.
func (s *BaseCqlListener) EnterCreateAggregateStatement(ctx *CreateAggregateStatementContext) {}

// ExitCreateAggregateStatement is called when production createAggregateStatement is exited.
func (s *BaseCqlListener) ExitCreateAggregateStatement(ctx *CreateAggregateStatementContext) {}

// EnterDropAggregateStatement is called when production dropAggregateStatement is entered.
func (s *BaseCqlListener) EnterDropAggregateStatement(ctx *DropAggregateStatementContext) {}

// ExitDropAggregateStatement is called when production dropAggregateStatement is exited.
func (s *BaseCqlListener) ExitDropAggregateStatement(ctx *DropAggregateStatementContext) {}

// EnterCreateFunctionStatement is called when production createFunctionStatement is entered.
func (s *BaseCqlListener) EnterCreateFunctionStatement(ctx *CreateFunctionStatementContext) {}

// ExitCreateFunctionStatement is called when production createFunctionStatement is exited.
func (s *BaseCqlListener) ExitCreateFunctionStatement(ctx *CreateFunctionStatementContext) {}

// EnterDropFunctionStatement is called when production dropFunctionStatement is entered.
func (s *BaseCqlListener) EnterDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// ExitDropFunctionStatement is called when production dropFunctionStatement is exited.
func (s *BaseCqlListener) ExitDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// EnterCreateKeyspaceStatement is called when production createKeyspaceStatement is entered.
func (s *BaseCqlListener) EnterCreateKeyspaceStatement(ctx *CreateKeyspaceStatementContext) {}

// ExitCreateKeyspaceStatement is called when production createKeyspaceStatement is exited.
func (s *BaseCqlListener) ExitCreateKeyspaceStatement(ctx *CreateKeyspaceStatementContext) {}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseCqlListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseCqlListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterTableDefinition is called when production tableDefinition is entered.
func (s *BaseCqlListener) EnterTableDefinition(ctx *TableDefinitionContext) {}

// ExitTableDefinition is called when production tableDefinition is exited.
func (s *BaseCqlListener) ExitTableDefinition(ctx *TableDefinitionContext) {}

// EnterTableColumns is called when production tableColumns is entered.
func (s *BaseCqlListener) EnterTableColumns(ctx *TableColumnsContext) {}

// ExitTableColumns is called when production tableColumns is exited.
func (s *BaseCqlListener) ExitTableColumns(ctx *TableColumnsContext) {}

// EnterColumnMask is called when production columnMask is entered.
func (s *BaseCqlListener) EnterColumnMask(ctx *ColumnMaskContext) {}

// ExitColumnMask is called when production columnMask is exited.
func (s *BaseCqlListener) ExitColumnMask(ctx *ColumnMaskContext) {}

// EnterColumnMaskArguments is called when production columnMaskArguments is entered.
func (s *BaseCqlListener) EnterColumnMaskArguments(ctx *ColumnMaskArgumentsContext) {}

// ExitColumnMaskArguments is called when production columnMaskArguments is exited.
func (s *BaseCqlListener) ExitColumnMaskArguments(ctx *ColumnMaskArgumentsContext) {}

// EnterTablePartitionKey is called when production tablePartitionKey is entered.
func (s *BaseCqlListener) EnterTablePartitionKey(ctx *TablePartitionKeyContext) {}

// ExitTablePartitionKey is called when production tablePartitionKey is exited.
func (s *BaseCqlListener) ExitTablePartitionKey(ctx *TablePartitionKeyContext) {}

// EnterTableProperty is called when production tableProperty is entered.
func (s *BaseCqlListener) EnterTableProperty(ctx *TablePropertyContext) {}

// ExitTableProperty is called when production tableProperty is exited.
func (s *BaseCqlListener) ExitTableProperty(ctx *TablePropertyContext) {}

// EnterTableClusteringOrder is called when production tableClusteringOrder is entered.
func (s *BaseCqlListener) EnterTableClusteringOrder(ctx *TableClusteringOrderContext) {}

// ExitTableClusteringOrder is called when production tableClusteringOrder is exited.
func (s *BaseCqlListener) ExitTableClusteringOrder(ctx *TableClusteringOrderContext) {}

// EnterCreateTypeStatement is called when production createTypeStatement is entered.
func (s *BaseCqlListener) EnterCreateTypeStatement(ctx *CreateTypeStatementContext) {}

// ExitCreateTypeStatement is called when production createTypeStatement is exited.
func (s *BaseCqlListener) ExitCreateTypeStatement(ctx *CreateTypeStatementContext) {}

// EnterTypeColumns is called when production typeColumns is entered.
func (s *BaseCqlListener) EnterTypeColumns(ctx *TypeColumnsContext) {}

// ExitTypeColumns is called when production typeColumns is exited.
func (s *BaseCqlListener) ExitTypeColumns(ctx *TypeColumnsContext) {}

// EnterCreateIndexStatement is called when production createIndexStatement is entered.
func (s *BaseCqlListener) EnterCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// ExitCreateIndexStatement is called when production createIndexStatement is exited.
func (s *BaseCqlListener) ExitCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// EnterIndexIdent is called when production indexIdent is entered.
func (s *BaseCqlListener) EnterIndexIdent(ctx *IndexIdentContext) {}

// ExitIndexIdent is called when production indexIdent is exited.
func (s *BaseCqlListener) ExitIndexIdent(ctx *IndexIdentContext) {}

// EnterCreateMaterializedViewStatement is called when production createMaterializedViewStatement is entered.
func (s *BaseCqlListener) EnterCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// ExitCreateMaterializedViewStatement is called when production createMaterializedViewStatement is exited.
func (s *BaseCqlListener) ExitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// EnterViewPrimaryKey is called when production viewPrimaryKey is entered.
func (s *BaseCqlListener) EnterViewPrimaryKey(ctx *ViewPrimaryKeyContext) {}

// ExitViewPrimaryKey is called when production viewPrimaryKey is exited.
func (s *BaseCqlListener) ExitViewPrimaryKey(ctx *ViewPrimaryKeyContext) {}

// EnterViewPartitionKey is called when production viewPartitionKey is entered.
func (s *BaseCqlListener) EnterViewPartitionKey(ctx *ViewPartitionKeyContext) {}

// ExitViewPartitionKey is called when production viewPartitionKey is exited.
func (s *BaseCqlListener) ExitViewPartitionKey(ctx *ViewPartitionKeyContext) {}

// EnterViewProperty is called when production viewProperty is entered.
func (s *BaseCqlListener) EnterViewProperty(ctx *ViewPropertyContext) {}

// ExitViewProperty is called when production viewProperty is exited.
func (s *BaseCqlListener) ExitViewProperty(ctx *ViewPropertyContext) {}

// EnterViewClusteringOrder is called when production viewClusteringOrder is entered.
func (s *BaseCqlListener) EnterViewClusteringOrder(ctx *ViewClusteringOrderContext) {}

// ExitViewClusteringOrder is called when production viewClusteringOrder is exited.
func (s *BaseCqlListener) ExitViewClusteringOrder(ctx *ViewClusteringOrderContext) {}

// EnterCreateTriggerStatement is called when production createTriggerStatement is entered.
func (s *BaseCqlListener) EnterCreateTriggerStatement(ctx *CreateTriggerStatementContext) {}

// ExitCreateTriggerStatement is called when production createTriggerStatement is exited.
func (s *BaseCqlListener) ExitCreateTriggerStatement(ctx *CreateTriggerStatementContext) {}

// EnterDropTriggerStatement is called when production dropTriggerStatement is entered.
func (s *BaseCqlListener) EnterDropTriggerStatement(ctx *DropTriggerStatementContext) {}

// ExitDropTriggerStatement is called when production dropTriggerStatement is exited.
func (s *BaseCqlListener) ExitDropTriggerStatement(ctx *DropTriggerStatementContext) {}

// EnterAlterKeyspaceStatement is called when production alterKeyspaceStatement is entered.
func (s *BaseCqlListener) EnterAlterKeyspaceStatement(ctx *AlterKeyspaceStatementContext) {}

// ExitAlterKeyspaceStatement is called when production alterKeyspaceStatement is exited.
func (s *BaseCqlListener) ExitAlterKeyspaceStatement(ctx *AlterKeyspaceStatementContext) {}

// EnterAlterTableStatement is called when production alterTableStatement is entered.
func (s *BaseCqlListener) EnterAlterTableStatement(ctx *AlterTableStatementContext) {}

// ExitAlterTableStatement is called when production alterTableStatement is exited.
func (s *BaseCqlListener) ExitAlterTableStatement(ctx *AlterTableStatementContext) {}

// EnterIsStaticColumn is called when production isStaticColumn is entered.
func (s *BaseCqlListener) EnterIsStaticColumn(ctx *IsStaticColumnContext) {}

// ExitIsStaticColumn is called when production isStaticColumn is exited.
func (s *BaseCqlListener) ExitIsStaticColumn(ctx *IsStaticColumnContext) {}

// EnterAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is entered.
func (s *BaseCqlListener) EnterAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// ExitAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is exited.
func (s *BaseCqlListener) ExitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// EnterAlterTypeStatement is called when production alterTypeStatement is entered.
func (s *BaseCqlListener) EnterAlterTypeStatement(ctx *AlterTypeStatementContext) {}

// ExitAlterTypeStatement is called when production alterTypeStatement is exited.
func (s *BaseCqlListener) ExitAlterTypeStatement(ctx *AlterTypeStatementContext) {}

// EnterDropKeyspaceStatement is called when production dropKeyspaceStatement is entered.
func (s *BaseCqlListener) EnterDropKeyspaceStatement(ctx *DropKeyspaceStatementContext) {}

// ExitDropKeyspaceStatement is called when production dropKeyspaceStatement is exited.
func (s *BaseCqlListener) ExitDropKeyspaceStatement(ctx *DropKeyspaceStatementContext) {}

// EnterDropTableStatement is called when production dropTableStatement is entered.
func (s *BaseCqlListener) EnterDropTableStatement(ctx *DropTableStatementContext) {}

// ExitDropTableStatement is called when production dropTableStatement is exited.
func (s *BaseCqlListener) ExitDropTableStatement(ctx *DropTableStatementContext) {}

// EnterDropTypeStatement is called when production dropTypeStatement is entered.
func (s *BaseCqlListener) EnterDropTypeStatement(ctx *DropTypeStatementContext) {}

// ExitDropTypeStatement is called when production dropTypeStatement is exited.
func (s *BaseCqlListener) ExitDropTypeStatement(ctx *DropTypeStatementContext) {}

// EnterDropIndexStatement is called when production dropIndexStatement is entered.
func (s *BaseCqlListener) EnterDropIndexStatement(ctx *DropIndexStatementContext) {}

// ExitDropIndexStatement is called when production dropIndexStatement is exited.
func (s *BaseCqlListener) ExitDropIndexStatement(ctx *DropIndexStatementContext) {}

// EnterDropMaterializedViewStatement is called when production dropMaterializedViewStatement is entered.
func (s *BaseCqlListener) EnterDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// ExitDropMaterializedViewStatement is called when production dropMaterializedViewStatement is exited.
func (s *BaseCqlListener) ExitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// EnterTruncateStatement is called when production truncateStatement is entered.
func (s *BaseCqlListener) EnterTruncateStatement(ctx *TruncateStatementContext) {}

// ExitTruncateStatement is called when production truncateStatement is exited.
func (s *BaseCqlListener) ExitTruncateStatement(ctx *TruncateStatementContext) {}

// EnterGrantPermissionsStatement is called when production grantPermissionsStatement is entered.
func (s *BaseCqlListener) EnterGrantPermissionsStatement(ctx *GrantPermissionsStatementContext) {}

// ExitGrantPermissionsStatement is called when production grantPermissionsStatement is exited.
func (s *BaseCqlListener) ExitGrantPermissionsStatement(ctx *GrantPermissionsStatementContext) {}

// EnterRevokePermissionsStatement is called when production revokePermissionsStatement is entered.
func (s *BaseCqlListener) EnterRevokePermissionsStatement(ctx *RevokePermissionsStatementContext) {}

// ExitRevokePermissionsStatement is called when production revokePermissionsStatement is exited.
func (s *BaseCqlListener) ExitRevokePermissionsStatement(ctx *RevokePermissionsStatementContext) {}

// EnterGrantRoleStatement is called when production grantRoleStatement is entered.
func (s *BaseCqlListener) EnterGrantRoleStatement(ctx *GrantRoleStatementContext) {}

// ExitGrantRoleStatement is called when production grantRoleStatement is exited.
func (s *BaseCqlListener) ExitGrantRoleStatement(ctx *GrantRoleStatementContext) {}

// EnterRevokeRoleStatement is called when production revokeRoleStatement is entered.
func (s *BaseCqlListener) EnterRevokeRoleStatement(ctx *RevokeRoleStatementContext) {}

// ExitRevokeRoleStatement is called when production revokeRoleStatement is exited.
func (s *BaseCqlListener) ExitRevokeRoleStatement(ctx *RevokeRoleStatementContext) {}

// EnterListPermissionsStatement is called when production listPermissionsStatement is entered.
func (s *BaseCqlListener) EnterListPermissionsStatement(ctx *ListPermissionsStatementContext) {}

// ExitListPermissionsStatement is called when production listPermissionsStatement is exited.
func (s *BaseCqlListener) ExitListPermissionsStatement(ctx *ListPermissionsStatementContext) {}

// EnterPermission is called when production permission is entered.
func (s *BaseCqlListener) EnterPermission(ctx *PermissionContext) {}

// ExitPermission is called when production permission is exited.
func (s *BaseCqlListener) ExitPermission(ctx *PermissionContext) {}

// EnterPermissionOrAll is called when production permissionOrAll is entered.
func (s *BaseCqlListener) EnterPermissionOrAll(ctx *PermissionOrAllContext) {}

// ExitPermissionOrAll is called when production permissionOrAll is exited.
func (s *BaseCqlListener) ExitPermissionOrAll(ctx *PermissionOrAllContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseCqlListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseCqlListener) ExitResource(ctx *ResourceContext) {}

// EnterDataResource is called when production dataResource is entered.
func (s *BaseCqlListener) EnterDataResource(ctx *DataResourceContext) {}

// ExitDataResource is called when production dataResource is exited.
func (s *BaseCqlListener) ExitDataResource(ctx *DataResourceContext) {}

// EnterJmxResource is called when production jmxResource is entered.
func (s *BaseCqlListener) EnterJmxResource(ctx *JmxResourceContext) {}

// ExitJmxResource is called when production jmxResource is exited.
func (s *BaseCqlListener) ExitJmxResource(ctx *JmxResourceContext) {}

// EnterRoleResource is called when production roleResource is entered.
func (s *BaseCqlListener) EnterRoleResource(ctx *RoleResourceContext) {}

// ExitRoleResource is called when production roleResource is exited.
func (s *BaseCqlListener) ExitRoleResource(ctx *RoleResourceContext) {}

// EnterFunctionResource is called when production functionResource is entered.
func (s *BaseCqlListener) EnterFunctionResource(ctx *FunctionResourceContext) {}

// ExitFunctionResource is called when production functionResource is exited.
func (s *BaseCqlListener) ExitFunctionResource(ctx *FunctionResourceContext) {}

// EnterCreateUserStatement is called when production createUserStatement is entered.
func (s *BaseCqlListener) EnterCreateUserStatement(ctx *CreateUserStatementContext) {}

// ExitCreateUserStatement is called when production createUserStatement is exited.
func (s *BaseCqlListener) ExitCreateUserStatement(ctx *CreateUserStatementContext) {}

// EnterAlterUserStatement is called when production alterUserStatement is entered.
func (s *BaseCqlListener) EnterAlterUserStatement(ctx *AlterUserStatementContext) {}

// ExitAlterUserStatement is called when production alterUserStatement is exited.
func (s *BaseCqlListener) ExitAlterUserStatement(ctx *AlterUserStatementContext) {}

// EnterDropUserStatement is called when production dropUserStatement is entered.
func (s *BaseCqlListener) EnterDropUserStatement(ctx *DropUserStatementContext) {}

// ExitDropUserStatement is called when production dropUserStatement is exited.
func (s *BaseCqlListener) ExitDropUserStatement(ctx *DropUserStatementContext) {}

// EnterAddIdentityStatement is called when production addIdentityStatement is entered.
func (s *BaseCqlListener) EnterAddIdentityStatement(ctx *AddIdentityStatementContext) {}

// ExitAddIdentityStatement is called when production addIdentityStatement is exited.
func (s *BaseCqlListener) ExitAddIdentityStatement(ctx *AddIdentityStatementContext) {}

// EnterDropIdentityStatement is called when production dropIdentityStatement is entered.
func (s *BaseCqlListener) EnterDropIdentityStatement(ctx *DropIdentityStatementContext) {}

// ExitDropIdentityStatement is called when production dropIdentityStatement is exited.
func (s *BaseCqlListener) ExitDropIdentityStatement(ctx *DropIdentityStatementContext) {}

// EnterListUsersStatement is called when production listUsersStatement is entered.
func (s *BaseCqlListener) EnterListUsersStatement(ctx *ListUsersStatementContext) {}

// ExitListUsersStatement is called when production listUsersStatement is exited.
func (s *BaseCqlListener) ExitListUsersStatement(ctx *ListUsersStatementContext) {}

// EnterCreateRoleStatement is called when production createRoleStatement is entered.
func (s *BaseCqlListener) EnterCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// ExitCreateRoleStatement is called when production createRoleStatement is exited.
func (s *BaseCqlListener) ExitCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// EnterAlterRoleStatement is called when production alterRoleStatement is entered.
func (s *BaseCqlListener) EnterAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// ExitAlterRoleStatement is called when production alterRoleStatement is exited.
func (s *BaseCqlListener) ExitAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// EnterDropRoleStatement is called when production dropRoleStatement is entered.
func (s *BaseCqlListener) EnterDropRoleStatement(ctx *DropRoleStatementContext) {}

// ExitDropRoleStatement is called when production dropRoleStatement is exited.
func (s *BaseCqlListener) ExitDropRoleStatement(ctx *DropRoleStatementContext) {}

// EnterListRolesStatement is called when production listRolesStatement is entered.
func (s *BaseCqlListener) EnterListRolesStatement(ctx *ListRolesStatementContext) {}

// ExitListRolesStatement is called when production listRolesStatement is exited.
func (s *BaseCqlListener) ExitListRolesStatement(ctx *ListRolesStatementContext) {}

// EnterRoleOptions is called when production roleOptions is entered.
func (s *BaseCqlListener) EnterRoleOptions(ctx *RoleOptionsContext) {}

// ExitRoleOptions is called when production roleOptions is exited.
func (s *BaseCqlListener) ExitRoleOptions(ctx *RoleOptionsContext) {}

// EnterRoleOption is called when production roleOption is entered.
func (s *BaseCqlListener) EnterRoleOption(ctx *RoleOptionContext) {}

// ExitRoleOption is called when production roleOption is exited.
func (s *BaseCqlListener) ExitRoleOption(ctx *RoleOptionContext) {}

// EnterDcPermission is called when production dcPermission is entered.
func (s *BaseCqlListener) EnterDcPermission(ctx *DcPermissionContext) {}

// ExitDcPermission is called when production dcPermission is exited.
func (s *BaseCqlListener) ExitDcPermission(ctx *DcPermissionContext) {}

// EnterCidrPermission is called when production cidrPermission is entered.
func (s *BaseCqlListener) EnterCidrPermission(ctx *CidrPermissionContext) {}

// ExitCidrPermission is called when production cidrPermission is exited.
func (s *BaseCqlListener) ExitCidrPermission(ctx *CidrPermissionContext) {}

// EnterUserPassword is called when production userPassword is entered.
func (s *BaseCqlListener) EnterUserPassword(ctx *UserPasswordContext) {}

// ExitUserPassword is called when production userPassword is exited.
func (s *BaseCqlListener) ExitUserPassword(ctx *UserPasswordContext) {}

// EnterDescribeStatement is called when production describeStatement is entered.
func (s *BaseCqlListener) EnterDescribeStatement(ctx *DescribeStatementContext) {}

// ExitDescribeStatement is called when production describeStatement is exited.
func (s *BaseCqlListener) ExitDescribeStatement(ctx *DescribeStatementContext) {}

// EnterCident is called when production cident is entered.
func (s *BaseCqlListener) EnterCident(ctx *CidentContext) {}

// ExitCident is called when production cident is exited.
func (s *BaseCqlListener) ExitCident(ctx *CidentContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseCqlListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseCqlListener) ExitIdent(ctx *IdentContext) {}

// EnterFident is called when production fident is entered.
func (s *BaseCqlListener) EnterFident(ctx *FidentContext) {}

// ExitFident is called when production fident is exited.
func (s *BaseCqlListener) ExitFident(ctx *FidentContext) {}

// EnterNoncol_ident is called when production noncol_ident is entered.
func (s *BaseCqlListener) EnterNoncol_ident(ctx *Noncol_identContext) {}

// ExitNoncol_ident is called when production noncol_ident is exited.
func (s *BaseCqlListener) ExitNoncol_ident(ctx *Noncol_identContext) {}

// EnterKeyspaceName is called when production keyspaceName is entered.
func (s *BaseCqlListener) EnterKeyspaceName(ctx *KeyspaceNameContext) {}

// ExitKeyspaceName is called when production keyspaceName is exited.
func (s *BaseCqlListener) ExitKeyspaceName(ctx *KeyspaceNameContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseCqlListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseCqlListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterColumnFamilyName is called when production columnFamilyName is entered.
func (s *BaseCqlListener) EnterColumnFamilyName(ctx *ColumnFamilyNameContext) {}

// ExitColumnFamilyName is called when production columnFamilyName is exited.
func (s *BaseCqlListener) ExitColumnFamilyName(ctx *ColumnFamilyNameContext) {}

// EnterUserTypeName is called when production userTypeName is entered.
func (s *BaseCqlListener) EnterUserTypeName(ctx *UserTypeNameContext) {}

// ExitUserTypeName is called when production userTypeName is exited.
func (s *BaseCqlListener) ExitUserTypeName(ctx *UserTypeNameContext) {}

// EnterUserOrRoleName is called when production userOrRoleName is entered.
func (s *BaseCqlListener) EnterUserOrRoleName(ctx *UserOrRoleNameContext) {}

// ExitUserOrRoleName is called when production userOrRoleName is exited.
func (s *BaseCqlListener) ExitUserOrRoleName(ctx *UserOrRoleNameContext) {}

// EnterKsName is called when production ksName is entered.
func (s *BaseCqlListener) EnterKsName(ctx *KsNameContext) {}

// ExitKsName is called when production ksName is exited.
func (s *BaseCqlListener) ExitKsName(ctx *KsNameContext) {}

// EnterCfName is called when production cfName is entered.
func (s *BaseCqlListener) EnterCfName(ctx *CfNameContext) {}

// ExitCfName is called when production cfName is exited.
func (s *BaseCqlListener) ExitCfName(ctx *CfNameContext) {}

// EnterIdxName is called when production idxName is entered.
func (s *BaseCqlListener) EnterIdxName(ctx *IdxNameContext) {}

// ExitIdxName is called when production idxName is exited.
func (s *BaseCqlListener) ExitIdxName(ctx *IdxNameContext) {}

// EnterRoleName is called when production roleName is entered.
func (s *BaseCqlListener) EnterRoleName(ctx *RoleNameContext) {}

// ExitRoleName is called when production roleName is exited.
func (s *BaseCqlListener) ExitRoleName(ctx *RoleNameContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCqlListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCqlListener) ExitConstant(ctx *ConstantContext) {}

// EnterFullMapLiteral is called when production fullMapLiteral is entered.
func (s *BaseCqlListener) EnterFullMapLiteral(ctx *FullMapLiteralContext) {}

// ExitFullMapLiteral is called when production fullMapLiteral is exited.
func (s *BaseCqlListener) ExitFullMapLiteral(ctx *FullMapLiteralContext) {}

// EnterSetOrMapLiteral is called when production setOrMapLiteral is entered.
func (s *BaseCqlListener) EnterSetOrMapLiteral(ctx *SetOrMapLiteralContext) {}

// ExitSetOrMapLiteral is called when production setOrMapLiteral is exited.
func (s *BaseCqlListener) ExitSetOrMapLiteral(ctx *SetOrMapLiteralContext) {}

// EnterSetLiteral is called when production setLiteral is entered.
func (s *BaseCqlListener) EnterSetLiteral(ctx *SetLiteralContext) {}

// ExitSetLiteral is called when production setLiteral is exited.
func (s *BaseCqlListener) ExitSetLiteral(ctx *SetLiteralContext) {}

// EnterMapLiteral is called when production mapLiteral is entered.
func (s *BaseCqlListener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production mapLiteral is exited.
func (s *BaseCqlListener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterCollectionLiteral is called when production collectionLiteral is entered.
func (s *BaseCqlListener) EnterCollectionLiteral(ctx *CollectionLiteralContext) {}

// ExitCollectionLiteral is called when production collectionLiteral is exited.
func (s *BaseCqlListener) ExitCollectionLiteral(ctx *CollectionLiteralContext) {}

// EnterListLiteral is called when production listLiteral is entered.
func (s *BaseCqlListener) EnterListLiteral(ctx *ListLiteralContext) {}

// ExitListLiteral is called when production listLiteral is exited.
func (s *BaseCqlListener) ExitListLiteral(ctx *ListLiteralContext) {}

// EnterUsertypeLiteral is called when production usertypeLiteral is entered.
func (s *BaseCqlListener) EnterUsertypeLiteral(ctx *UsertypeLiteralContext) {}

// ExitUsertypeLiteral is called when production usertypeLiteral is exited.
func (s *BaseCqlListener) ExitUsertypeLiteral(ctx *UsertypeLiteralContext) {}

// EnterTupleLiteral is called when production tupleLiteral is entered.
func (s *BaseCqlListener) EnterTupleLiteral(ctx *TupleLiteralContext) {}

// ExitTupleLiteral is called when production tupleLiteral is exited.
func (s *BaseCqlListener) ExitTupleLiteral(ctx *TupleLiteralContext) {}

// EnterValue is called when production value is entered.
func (s *BaseCqlListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseCqlListener) ExitValue(ctx *ValueContext) {}

// EnterIntValue is called when production intValue is entered.
func (s *BaseCqlListener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production intValue is exited.
func (s *BaseCqlListener) ExitIntValue(ctx *IntValueContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseCqlListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseCqlListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterAllowedFunctionName is called when production allowedFunctionName is entered.
func (s *BaseCqlListener) EnterAllowedFunctionName(ctx *AllowedFunctionNameContext) {}

// ExitAllowedFunctionName is called when production allowedFunctionName is exited.
func (s *BaseCqlListener) ExitAllowedFunctionName(ctx *AllowedFunctionNameContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseCqlListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseCqlListener) ExitFunction(ctx *FunctionContext) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BaseCqlListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BaseCqlListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseCqlListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseCqlListener) ExitTerm(ctx *TermContext) {}

// EnterTermAddition is called when production termAddition is entered.
func (s *BaseCqlListener) EnterTermAddition(ctx *TermAdditionContext) {}

// ExitTermAddition is called when production termAddition is exited.
func (s *BaseCqlListener) ExitTermAddition(ctx *TermAdditionContext) {}

// EnterTermMultiplication is called when production termMultiplication is entered.
func (s *BaseCqlListener) EnterTermMultiplication(ctx *TermMultiplicationContext) {}

// ExitTermMultiplication is called when production termMultiplication is exited.
func (s *BaseCqlListener) ExitTermMultiplication(ctx *TermMultiplicationContext) {}

// EnterTermGroup is called when production termGroup is entered.
func (s *BaseCqlListener) EnterTermGroup(ctx *TermGroupContext) {}

// ExitTermGroup is called when production termGroup is exited.
func (s *BaseCqlListener) ExitTermGroup(ctx *TermGroupContext) {}

// EnterSimpleTerm is called when production simpleTerm is entered.
func (s *BaseCqlListener) EnterSimpleTerm(ctx *SimpleTermContext) {}

// ExitSimpleTerm is called when production simpleTerm is exited.
func (s *BaseCqlListener) ExitSimpleTerm(ctx *SimpleTermContext) {}

// EnterColumnOperation is called when production columnOperation is entered.
func (s *BaseCqlListener) EnterColumnOperation(ctx *ColumnOperationContext) {}

// ExitColumnOperation is called when production columnOperation is exited.
func (s *BaseCqlListener) ExitColumnOperation(ctx *ColumnOperationContext) {}

// EnterColumnOperationDifferentiator is called when production columnOperationDifferentiator is entered.
func (s *BaseCqlListener) EnterColumnOperationDifferentiator(ctx *ColumnOperationDifferentiatorContext) {
}

// ExitColumnOperationDifferentiator is called when production columnOperationDifferentiator is exited.
func (s *BaseCqlListener) ExitColumnOperationDifferentiator(ctx *ColumnOperationDifferentiatorContext) {
}

// EnterNormalColumnOperation is called when production normalColumnOperation is entered.
func (s *BaseCqlListener) EnterNormalColumnOperation(ctx *NormalColumnOperationContext) {}

// ExitNormalColumnOperation is called when production normalColumnOperation is exited.
func (s *BaseCqlListener) ExitNormalColumnOperation(ctx *NormalColumnOperationContext) {}

// EnterShorthandColumnOperation is called when production shorthandColumnOperation is entered.
func (s *BaseCqlListener) EnterShorthandColumnOperation(ctx *ShorthandColumnOperationContext) {}

// ExitShorthandColumnOperation is called when production shorthandColumnOperation is exited.
func (s *BaseCqlListener) ExitShorthandColumnOperation(ctx *ShorthandColumnOperationContext) {}

// EnterCollectionColumnOperation is called when production collectionColumnOperation is entered.
func (s *BaseCqlListener) EnterCollectionColumnOperation(ctx *CollectionColumnOperationContext) {}

// ExitCollectionColumnOperation is called when production collectionColumnOperation is exited.
func (s *BaseCqlListener) ExitCollectionColumnOperation(ctx *CollectionColumnOperationContext) {}

// EnterUdtColumnOperation is called when production udtColumnOperation is entered.
func (s *BaseCqlListener) EnterUdtColumnOperation(ctx *UdtColumnOperationContext) {}

// ExitUdtColumnOperation is called when production udtColumnOperation is exited.
func (s *BaseCqlListener) ExitUdtColumnOperation(ctx *UdtColumnOperationContext) {}

// EnterColumnCondition is called when production columnCondition is entered.
func (s *BaseCqlListener) EnterColumnCondition(ctx *ColumnConditionContext) {}

// ExitColumnCondition is called when production columnCondition is exited.
func (s *BaseCqlListener) ExitColumnCondition(ctx *ColumnConditionContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseCqlListener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseCqlListener) ExitProperties(ctx *PropertiesContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseCqlListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseCqlListener) ExitProperty(ctx *PropertyContext) {}

// EnterPropertyValue is called when production propertyValue is entered.
func (s *BaseCqlListener) EnterPropertyValue(ctx *PropertyValueContext) {}

// ExitPropertyValue is called when production propertyValue is exited.
func (s *BaseCqlListener) ExitPropertyValue(ctx *PropertyValueContext) {}

// EnterRelationType is called when production relationType is entered.
func (s *BaseCqlListener) EnterRelationType(ctx *RelationTypeContext) {}

// ExitRelationType is called when production relationType is exited.
func (s *BaseCqlListener) ExitRelationType(ctx *RelationTypeContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseCqlListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseCqlListener) ExitRelation(ctx *RelationContext) {}

// EnterContainsOperator is called when production containsOperator is entered.
func (s *BaseCqlListener) EnterContainsOperator(ctx *ContainsOperatorContext) {}

// ExitContainsOperator is called when production containsOperator is exited.
func (s *BaseCqlListener) ExitContainsOperator(ctx *ContainsOperatorContext) {}

// EnterInMarker is called when production inMarker is entered.
func (s *BaseCqlListener) EnterInMarker(ctx *InMarkerContext) {}

// ExitInMarker is called when production inMarker is exited.
func (s *BaseCqlListener) ExitInMarker(ctx *InMarkerContext) {}

// EnterTupleOfIdentifiers is called when production tupleOfIdentifiers is entered.
func (s *BaseCqlListener) EnterTupleOfIdentifiers(ctx *TupleOfIdentifiersContext) {}

// ExitTupleOfIdentifiers is called when production tupleOfIdentifiers is exited.
func (s *BaseCqlListener) ExitTupleOfIdentifiers(ctx *TupleOfIdentifiersContext) {}

// EnterSingleColumnInValues is called when production singleColumnInValues is entered.
func (s *BaseCqlListener) EnterSingleColumnInValues(ctx *SingleColumnInValuesContext) {}

// ExitSingleColumnInValues is called when production singleColumnInValues is exited.
func (s *BaseCqlListener) ExitSingleColumnInValues(ctx *SingleColumnInValuesContext) {}

// EnterTupleOfTupleLiterals is called when production tupleOfTupleLiterals is entered.
func (s *BaseCqlListener) EnterTupleOfTupleLiterals(ctx *TupleOfTupleLiteralsContext) {}

// ExitTupleOfTupleLiterals is called when production tupleOfTupleLiterals is exited.
func (s *BaseCqlListener) ExitTupleOfTupleLiterals(ctx *TupleOfTupleLiteralsContext) {}

// EnterMarkerForTuple is called when production markerForTuple is entered.
func (s *BaseCqlListener) EnterMarkerForTuple(ctx *MarkerForTupleContext) {}

// ExitMarkerForTuple is called when production markerForTuple is exited.
func (s *BaseCqlListener) ExitMarkerForTuple(ctx *MarkerForTupleContext) {}

// EnterTupleOfMarkersForTuples is called when production tupleOfMarkersForTuples is entered.
func (s *BaseCqlListener) EnterTupleOfMarkersForTuples(ctx *TupleOfMarkersForTuplesContext) {}

// ExitTupleOfMarkersForTuples is called when production tupleOfMarkersForTuples is exited.
func (s *BaseCqlListener) ExitTupleOfMarkersForTuples(ctx *TupleOfMarkersForTuplesContext) {}

// EnterInMarkerForTuple is called when production inMarkerForTuple is entered.
func (s *BaseCqlListener) EnterInMarkerForTuple(ctx *InMarkerForTupleContext) {}

// ExitInMarkerForTuple is called when production inMarkerForTuple is exited.
func (s *BaseCqlListener) ExitInMarkerForTuple(ctx *InMarkerForTupleContext) {}

// EnterComparatorType is called when production comparatorType is entered.
func (s *BaseCqlListener) EnterComparatorType(ctx *ComparatorTypeContext) {}

// ExitComparatorType is called when production comparatorType is exited.
func (s *BaseCqlListener) ExitComparatorType(ctx *ComparatorTypeContext) {}

// EnterNative_type is called when production native_type is entered.
func (s *BaseCqlListener) EnterNative_type(ctx *Native_typeContext) {}

// ExitNative_type is called when production native_type is exited.
func (s *BaseCqlListener) ExitNative_type(ctx *Native_typeContext) {}

// EnterCollection_type is called when production collection_type is entered.
func (s *BaseCqlListener) EnterCollection_type(ctx *Collection_typeContext) {}

// ExitCollection_type is called when production collection_type is exited.
func (s *BaseCqlListener) ExitCollection_type(ctx *Collection_typeContext) {}

// EnterTuple_type is called when production tuple_type is entered.
func (s *BaseCqlListener) EnterTuple_type(ctx *Tuple_typeContext) {}

// ExitTuple_type is called when production tuple_type is exited.
func (s *BaseCqlListener) ExitTuple_type(ctx *Tuple_typeContext) {}

// EnterVector_type is called when production vector_type is entered.
func (s *BaseCqlListener) EnterVector_type(ctx *Vector_typeContext) {}

// ExitVector_type is called when production vector_type is exited.
func (s *BaseCqlListener) ExitVector_type(ctx *Vector_typeContext) {}

// EnterUsername is called when production username is entered.
func (s *BaseCqlListener) EnterUsername(ctx *UsernameContext) {}

// ExitUsername is called when production username is exited.
func (s *BaseCqlListener) ExitUsername(ctx *UsernameContext) {}

// EnterIdentity is called when production identity is entered.
func (s *BaseCqlListener) EnterIdentity(ctx *IdentityContext) {}

// ExitIdentity is called when production identity is exited.
func (s *BaseCqlListener) ExitIdentity(ctx *IdentityContext) {}

// EnterMbean is called when production mbean is entered.
func (s *BaseCqlListener) EnterMbean(ctx *MbeanContext) {}

// ExitMbean is called when production mbean is exited.
func (s *BaseCqlListener) ExitMbean(ctx *MbeanContext) {}

// EnterNon_type_ident is called when production non_type_ident is entered.
func (s *BaseCqlListener) EnterNon_type_ident(ctx *Non_type_identContext) {}

// ExitNon_type_ident is called when production non_type_ident is exited.
func (s *BaseCqlListener) ExitNon_type_ident(ctx *Non_type_identContext) {}

// EnterUnreserved_keyword is called when production unreserved_keyword is entered.
func (s *BaseCqlListener) EnterUnreserved_keyword(ctx *Unreserved_keywordContext) {}

// ExitUnreserved_keyword is called when production unreserved_keyword is exited.
func (s *BaseCqlListener) ExitUnreserved_keyword(ctx *Unreserved_keywordContext) {}

// EnterUnreserved_function_keyword is called when production unreserved_function_keyword is entered.
func (s *BaseCqlListener) EnterUnreserved_function_keyword(ctx *Unreserved_function_keywordContext) {}

// ExitUnreserved_function_keyword is called when production unreserved_function_keyword is exited.
func (s *BaseCqlListener) ExitUnreserved_function_keyword(ctx *Unreserved_function_keywordContext) {}

// EnterBasic_unreserved_keyword is called when production basic_unreserved_keyword is entered.
func (s *BaseCqlListener) EnterBasic_unreserved_keyword(ctx *Basic_unreserved_keywordContext) {}

// ExitBasic_unreserved_keyword is called when production basic_unreserved_keyword is exited.
func (s *BaseCqlListener) ExitBasic_unreserved_keyword(ctx *Basic_unreserved_keywordContext) {}
