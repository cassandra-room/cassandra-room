// Code generated from /Users/jungwook/cassandra-room/internal/cql/parser/Cql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Cql
import "github.com/antlr4-go/antlr/v4"

// CqlListener is a complete listener for a parse tree produced by CqlParser.
type CqlListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterCqlStatements is called when entering the cqlStatements production.
	EnterCqlStatements(c *CqlStatementsContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterCqlStatement is called when entering the cqlStatement production.
	EnterCqlStatement(c *CqlStatementContext)

	// EnterUseStatement is called when entering the useStatement production.
	EnterUseStatement(c *UseStatementContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterSelectors is called when entering the selectors production.
	EnterSelectors(c *SelectorsContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterUnaliasedSelector is called when entering the unaliasedSelector production.
	EnterUnaliasedSelector(c *UnaliasedSelectorContext)

	// EnterSelectionAddition is called when entering the selectionAddition production.
	EnterSelectionAddition(c *SelectionAdditionContext)

	// EnterSelectionMultiplication is called when entering the selectionMultiplication production.
	EnterSelectionMultiplication(c *SelectionMultiplicationContext)

	// EnterSelectionGroup is called when entering the selectionGroup production.
	EnterSelectionGroup(c *SelectionGroupContext)

	// EnterSelectionGroupWithField is called when entering the selectionGroupWithField production.
	EnterSelectionGroupWithField(c *SelectionGroupWithFieldContext)

	// EnterSelectorModifier is called when entering the selectorModifier production.
	EnterSelectorModifier(c *SelectorModifierContext)

	// EnterFieldSelectorModifier is called when entering the fieldSelectorModifier production.
	EnterFieldSelectorModifier(c *FieldSelectorModifierContext)

	// EnterCollectionSubSelection is called when entering the collectionSubSelection production.
	EnterCollectionSubSelection(c *CollectionSubSelectionContext)

	// EnterSelectionGroupWithoutField is called when entering the selectionGroupWithoutField production.
	EnterSelectionGroupWithoutField(c *SelectionGroupWithoutFieldContext)

	// EnterSelectionTypeHint is called when entering the selectionTypeHint production.
	EnterSelectionTypeHint(c *SelectionTypeHintContext)

	// EnterSelectionList is called when entering the selectionList production.
	EnterSelectionList(c *SelectionListContext)

	// EnterSelectionMapOrSet is called when entering the selectionMapOrSet production.
	EnterSelectionMapOrSet(c *SelectionMapOrSetContext)

	// EnterSelectionMap is called when entering the selectionMap production.
	EnterSelectionMap(c *SelectionMapContext)

	// EnterSelectionSet is called when entering the selectionSet production.
	EnterSelectionSet(c *SelectionSetContext)

	// EnterSelectionTupleOrNestedSelector is called when entering the selectionTupleOrNestedSelector production.
	EnterSelectionTupleOrNestedSelector(c *SelectionTupleOrNestedSelectorContext)

	// EnterSimpleUnaliasedSelector is called when entering the simpleUnaliasedSelector production.
	EnterSimpleUnaliasedSelector(c *SimpleUnaliasedSelectorContext)

	// EnterSelectionFunction is called when entering the selectionFunction production.
	EnterSelectionFunction(c *SelectionFunctionContext)

	// EnterSelectionLiteral is called when entering the selectionLiteral production.
	EnterSelectionLiteral(c *SelectionLiteralContext)

	// EnterSelectionFunctionArgs is called when entering the selectionFunctionArgs production.
	EnterSelectionFunctionArgs(c *SelectionFunctionArgsContext)

	// EnterSident is called when entering the sident production.
	EnterSident(c *SidentContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterRelationOrExpression is called when entering the relationOrExpression production.
	EnterRelationOrExpression(c *RelationOrExpressionContext)

	// EnterCustomIndexExpression is called when entering the customIndexExpression production.
	EnterCustomIndexExpression(c *CustomIndexExpressionContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterNormalInsertStatement is called when entering the normalInsertStatement production.
	EnterNormalInsertStatement(c *NormalInsertStatementContext)

	// EnterJsonInsertStatement is called when entering the jsonInsertStatement production.
	EnterJsonInsertStatement(c *JsonInsertStatementContext)

	// EnterJsonValue is called when entering the jsonValue production.
	EnterJsonValue(c *JsonValueContext)

	// EnterUsingClause is called when entering the usingClause production.
	EnterUsingClause(c *UsingClauseContext)

	// EnterUsingClauseObjective is called when entering the usingClauseObjective production.
	EnterUsingClauseObjective(c *UsingClauseObjectiveContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterUpdateConditions is called when entering the updateConditions production.
	EnterUpdateConditions(c *UpdateConditionsContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterDeleteSelection is called when entering the deleteSelection production.
	EnterDeleteSelection(c *DeleteSelectionContext)

	// EnterDeleteOp is called when entering the deleteOp production.
	EnterDeleteOp(c *DeleteOpContext)

	// EnterUsingClauseDelete is called when entering the usingClauseDelete production.
	EnterUsingClauseDelete(c *UsingClauseDeleteContext)

	// EnterBatchStatement is called when entering the batchStatement production.
	EnterBatchStatement(c *BatchStatementContext)

	// EnterBatchStatementObjective is called when entering the batchStatementObjective production.
	EnterBatchStatementObjective(c *BatchStatementObjectiveContext)

	// EnterCreateAggregateStatement is called when entering the createAggregateStatement production.
	EnterCreateAggregateStatement(c *CreateAggregateStatementContext)

	// EnterDropAggregateStatement is called when entering the dropAggregateStatement production.
	EnterDropAggregateStatement(c *DropAggregateStatementContext)

	// EnterCreateFunctionStatement is called when entering the createFunctionStatement production.
	EnterCreateFunctionStatement(c *CreateFunctionStatementContext)

	// EnterDropFunctionStatement is called when entering the dropFunctionStatement production.
	EnterDropFunctionStatement(c *DropFunctionStatementContext)

	// EnterCreateKeyspaceStatement is called when entering the createKeyspaceStatement production.
	EnterCreateKeyspaceStatement(c *CreateKeyspaceStatementContext)

	// EnterCreateTableStatement is called when entering the createTableStatement production.
	EnterCreateTableStatement(c *CreateTableStatementContext)

	// EnterTableDefinition is called when entering the tableDefinition production.
	EnterTableDefinition(c *TableDefinitionContext)

	// EnterTableColumns is called when entering the tableColumns production.
	EnterTableColumns(c *TableColumnsContext)

	// EnterColumnMask is called when entering the columnMask production.
	EnterColumnMask(c *ColumnMaskContext)

	// EnterColumnMaskArguments is called when entering the columnMaskArguments production.
	EnterColumnMaskArguments(c *ColumnMaskArgumentsContext)

	// EnterTablePartitionKey is called when entering the tablePartitionKey production.
	EnterTablePartitionKey(c *TablePartitionKeyContext)

	// EnterTableProperty is called when entering the tableProperty production.
	EnterTableProperty(c *TablePropertyContext)

	// EnterTableClusteringOrder is called when entering the tableClusteringOrder production.
	EnterTableClusteringOrder(c *TableClusteringOrderContext)

	// EnterCreateTypeStatement is called when entering the createTypeStatement production.
	EnterCreateTypeStatement(c *CreateTypeStatementContext)

	// EnterTypeColumns is called when entering the typeColumns production.
	EnterTypeColumns(c *TypeColumnsContext)

	// EnterCreateIndexStatement is called when entering the createIndexStatement production.
	EnterCreateIndexStatement(c *CreateIndexStatementContext)

	// EnterIndexIdent is called when entering the indexIdent production.
	EnterIndexIdent(c *IndexIdentContext)

	// EnterCreateMaterializedViewStatement is called when entering the createMaterializedViewStatement production.
	EnterCreateMaterializedViewStatement(c *CreateMaterializedViewStatementContext)

	// EnterViewPrimaryKey is called when entering the viewPrimaryKey production.
	EnterViewPrimaryKey(c *ViewPrimaryKeyContext)

	// EnterViewPartitionKey is called when entering the viewPartitionKey production.
	EnterViewPartitionKey(c *ViewPartitionKeyContext)

	// EnterViewProperty is called when entering the viewProperty production.
	EnterViewProperty(c *ViewPropertyContext)

	// EnterViewClusteringOrder is called when entering the viewClusteringOrder production.
	EnterViewClusteringOrder(c *ViewClusteringOrderContext)

	// EnterCreateTriggerStatement is called when entering the createTriggerStatement production.
	EnterCreateTriggerStatement(c *CreateTriggerStatementContext)

	// EnterDropTriggerStatement is called when entering the dropTriggerStatement production.
	EnterDropTriggerStatement(c *DropTriggerStatementContext)

	// EnterAlterKeyspaceStatement is called when entering the alterKeyspaceStatement production.
	EnterAlterKeyspaceStatement(c *AlterKeyspaceStatementContext)

	// EnterAlterTableStatement is called when entering the alterTableStatement production.
	EnterAlterTableStatement(c *AlterTableStatementContext)

	// EnterIsStaticColumn is called when entering the isStaticColumn production.
	EnterIsStaticColumn(c *IsStaticColumnContext)

	// EnterAlterMaterializedViewStatement is called when entering the alterMaterializedViewStatement production.
	EnterAlterMaterializedViewStatement(c *AlterMaterializedViewStatementContext)

	// EnterAlterTypeStatement is called when entering the alterTypeStatement production.
	EnterAlterTypeStatement(c *AlterTypeStatementContext)

	// EnterDropKeyspaceStatement is called when entering the dropKeyspaceStatement production.
	EnterDropKeyspaceStatement(c *DropKeyspaceStatementContext)

	// EnterDropTableStatement is called when entering the dropTableStatement production.
	EnterDropTableStatement(c *DropTableStatementContext)

	// EnterDropTypeStatement is called when entering the dropTypeStatement production.
	EnterDropTypeStatement(c *DropTypeStatementContext)

	// EnterDropIndexStatement is called when entering the dropIndexStatement production.
	EnterDropIndexStatement(c *DropIndexStatementContext)

	// EnterDropMaterializedViewStatement is called when entering the dropMaterializedViewStatement production.
	EnterDropMaterializedViewStatement(c *DropMaterializedViewStatementContext)

	// EnterTruncateStatement is called when entering the truncateStatement production.
	EnterTruncateStatement(c *TruncateStatementContext)

	// EnterGrantPermissionsStatement is called when entering the grantPermissionsStatement production.
	EnterGrantPermissionsStatement(c *GrantPermissionsStatementContext)

	// EnterRevokePermissionsStatement is called when entering the revokePermissionsStatement production.
	EnterRevokePermissionsStatement(c *RevokePermissionsStatementContext)

	// EnterGrantRoleStatement is called when entering the grantRoleStatement production.
	EnterGrantRoleStatement(c *GrantRoleStatementContext)

	// EnterRevokeRoleStatement is called when entering the revokeRoleStatement production.
	EnterRevokeRoleStatement(c *RevokeRoleStatementContext)

	// EnterListPermissionsStatement is called when entering the listPermissionsStatement production.
	EnterListPermissionsStatement(c *ListPermissionsStatementContext)

	// EnterPermission is called when entering the permission production.
	EnterPermission(c *PermissionContext)

	// EnterPermissionOrAll is called when entering the permissionOrAll production.
	EnterPermissionOrAll(c *PermissionOrAllContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterDataResource is called when entering the dataResource production.
	EnterDataResource(c *DataResourceContext)

	// EnterJmxResource is called when entering the jmxResource production.
	EnterJmxResource(c *JmxResourceContext)

	// EnterRoleResource is called when entering the roleResource production.
	EnterRoleResource(c *RoleResourceContext)

	// EnterFunctionResource is called when entering the functionResource production.
	EnterFunctionResource(c *FunctionResourceContext)

	// EnterCreateUserStatement is called when entering the createUserStatement production.
	EnterCreateUserStatement(c *CreateUserStatementContext)

	// EnterAlterUserStatement is called when entering the alterUserStatement production.
	EnterAlterUserStatement(c *AlterUserStatementContext)

	// EnterDropUserStatement is called when entering the dropUserStatement production.
	EnterDropUserStatement(c *DropUserStatementContext)

	// EnterAddIdentityStatement is called when entering the addIdentityStatement production.
	EnterAddIdentityStatement(c *AddIdentityStatementContext)

	// EnterDropIdentityStatement is called when entering the dropIdentityStatement production.
	EnterDropIdentityStatement(c *DropIdentityStatementContext)

	// EnterListUsersStatement is called when entering the listUsersStatement production.
	EnterListUsersStatement(c *ListUsersStatementContext)

	// EnterCreateRoleStatement is called when entering the createRoleStatement production.
	EnterCreateRoleStatement(c *CreateRoleStatementContext)

	// EnterAlterRoleStatement is called when entering the alterRoleStatement production.
	EnterAlterRoleStatement(c *AlterRoleStatementContext)

	// EnterDropRoleStatement is called when entering the dropRoleStatement production.
	EnterDropRoleStatement(c *DropRoleStatementContext)

	// EnterListRolesStatement is called when entering the listRolesStatement production.
	EnterListRolesStatement(c *ListRolesStatementContext)

	// EnterRoleOptions is called when entering the roleOptions production.
	EnterRoleOptions(c *RoleOptionsContext)

	// EnterRoleOption is called when entering the roleOption production.
	EnterRoleOption(c *RoleOptionContext)

	// EnterDcPermission is called when entering the dcPermission production.
	EnterDcPermission(c *DcPermissionContext)

	// EnterCidrPermission is called when entering the cidrPermission production.
	EnterCidrPermission(c *CidrPermissionContext)

	// EnterUserPassword is called when entering the userPassword production.
	EnterUserPassword(c *UserPasswordContext)

	// EnterDescribeStatement is called when entering the describeStatement production.
	EnterDescribeStatement(c *DescribeStatementContext)

	// EnterCident is called when entering the cident production.
	EnterCident(c *CidentContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterFident is called when entering the fident production.
	EnterFident(c *FidentContext)

	// EnterNoncol_ident is called when entering the noncol_ident production.
	EnterNoncol_ident(c *Noncol_identContext)

	// EnterKeyspaceName is called when entering the keyspaceName production.
	EnterKeyspaceName(c *KeyspaceNameContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterColumnFamilyName is called when entering the columnFamilyName production.
	EnterColumnFamilyName(c *ColumnFamilyNameContext)

	// EnterUserTypeName is called when entering the userTypeName production.
	EnterUserTypeName(c *UserTypeNameContext)

	// EnterUserOrRoleName is called when entering the userOrRoleName production.
	EnterUserOrRoleName(c *UserOrRoleNameContext)

	// EnterKsName is called when entering the ksName production.
	EnterKsName(c *KsNameContext)

	// EnterCfName is called when entering the cfName production.
	EnterCfName(c *CfNameContext)

	// EnterIdxName is called when entering the idxName production.
	EnterIdxName(c *IdxNameContext)

	// EnterRoleName is called when entering the roleName production.
	EnterRoleName(c *RoleNameContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterFullMapLiteral is called when entering the fullMapLiteral production.
	EnterFullMapLiteral(c *FullMapLiteralContext)

	// EnterSetOrMapLiteral is called when entering the setOrMapLiteral production.
	EnterSetOrMapLiteral(c *SetOrMapLiteralContext)

	// EnterSetLiteral is called when entering the setLiteral production.
	EnterSetLiteral(c *SetLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterCollectionLiteral is called when entering the collectionLiteral production.
	EnterCollectionLiteral(c *CollectionLiteralContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterUsertypeLiteral is called when entering the usertypeLiteral production.
	EnterUsertypeLiteral(c *UsertypeLiteralContext)

	// EnterTupleLiteral is called when entering the tupleLiteral production.
	EnterTupleLiteral(c *TupleLiteralContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterIntValue is called when entering the intValue production.
	EnterIntValue(c *IntValueContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterAllowedFunctionName is called when entering the allowedFunctionName production.
	EnterAllowedFunctionName(c *AllowedFunctionNameContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFunctionArgs is called when entering the functionArgs production.
	EnterFunctionArgs(c *FunctionArgsContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterTermAddition is called when entering the termAddition production.
	EnterTermAddition(c *TermAdditionContext)

	// EnterTermMultiplication is called when entering the termMultiplication production.
	EnterTermMultiplication(c *TermMultiplicationContext)

	// EnterTermGroup is called when entering the termGroup production.
	EnterTermGroup(c *TermGroupContext)

	// EnterSimpleTerm is called when entering the simpleTerm production.
	EnterSimpleTerm(c *SimpleTermContext)

	// EnterColumnOperation is called when entering the columnOperation production.
	EnterColumnOperation(c *ColumnOperationContext)

	// EnterColumnOperationDifferentiator is called when entering the columnOperationDifferentiator production.
	EnterColumnOperationDifferentiator(c *ColumnOperationDifferentiatorContext)

	// EnterNormalColumnOperation is called when entering the normalColumnOperation production.
	EnterNormalColumnOperation(c *NormalColumnOperationContext)

	// EnterShorthandColumnOperation is called when entering the shorthandColumnOperation production.
	EnterShorthandColumnOperation(c *ShorthandColumnOperationContext)

	// EnterCollectionColumnOperation is called when entering the collectionColumnOperation production.
	EnterCollectionColumnOperation(c *CollectionColumnOperationContext)

	// EnterUdtColumnOperation is called when entering the udtColumnOperation production.
	EnterUdtColumnOperation(c *UdtColumnOperationContext)

	// EnterColumnCondition is called when entering the columnCondition production.
	EnterColumnCondition(c *ColumnConditionContext)

	// EnterProperties is called when entering the properties production.
	EnterProperties(c *PropertiesContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterPropertyValue is called when entering the propertyValue production.
	EnterPropertyValue(c *PropertyValueContext)

	// EnterRelationType is called when entering the relationType production.
	EnterRelationType(c *RelationTypeContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterContainsOperator is called when entering the containsOperator production.
	EnterContainsOperator(c *ContainsOperatorContext)

	// EnterInMarker is called when entering the inMarker production.
	EnterInMarker(c *InMarkerContext)

	// EnterTupleOfIdentifiers is called when entering the tupleOfIdentifiers production.
	EnterTupleOfIdentifiers(c *TupleOfIdentifiersContext)

	// EnterSingleColumnInValues is called when entering the singleColumnInValues production.
	EnterSingleColumnInValues(c *SingleColumnInValuesContext)

	// EnterTupleOfTupleLiterals is called when entering the tupleOfTupleLiterals production.
	EnterTupleOfTupleLiterals(c *TupleOfTupleLiteralsContext)

	// EnterMarkerForTuple is called when entering the markerForTuple production.
	EnterMarkerForTuple(c *MarkerForTupleContext)

	// EnterTupleOfMarkersForTuples is called when entering the tupleOfMarkersForTuples production.
	EnterTupleOfMarkersForTuples(c *TupleOfMarkersForTuplesContext)

	// EnterInMarkerForTuple is called when entering the inMarkerForTuple production.
	EnterInMarkerForTuple(c *InMarkerForTupleContext)

	// EnterComparatorType is called when entering the comparatorType production.
	EnterComparatorType(c *ComparatorTypeContext)

	// EnterNative_type is called when entering the native_type production.
	EnterNative_type(c *Native_typeContext)

	// EnterCollection_type is called when entering the collection_type production.
	EnterCollection_type(c *Collection_typeContext)

	// EnterTuple_type is called when entering the tuple_type production.
	EnterTuple_type(c *Tuple_typeContext)

	// EnterVector_type is called when entering the vector_type production.
	EnterVector_type(c *Vector_typeContext)

	// EnterUsername is called when entering the username production.
	EnterUsername(c *UsernameContext)

	// EnterIdentity is called when entering the identity production.
	EnterIdentity(c *IdentityContext)

	// EnterMbean is called when entering the mbean production.
	EnterMbean(c *MbeanContext)

	// EnterNon_type_ident is called when entering the non_type_ident production.
	EnterNon_type_ident(c *Non_type_identContext)

	// EnterUnreserved_keyword is called when entering the unreserved_keyword production.
	EnterUnreserved_keyword(c *Unreserved_keywordContext)

	// EnterUnreserved_function_keyword is called when entering the unreserved_function_keyword production.
	EnterUnreserved_function_keyword(c *Unreserved_function_keywordContext)

	// EnterBasic_unreserved_keyword is called when entering the basic_unreserved_keyword production.
	EnterBasic_unreserved_keyword(c *Basic_unreserved_keywordContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitCqlStatements is called when exiting the cqlStatements production.
	ExitCqlStatements(c *CqlStatementsContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitCqlStatement is called when exiting the cqlStatement production.
	ExitCqlStatement(c *CqlStatementContext)

	// ExitUseStatement is called when exiting the useStatement production.
	ExitUseStatement(c *UseStatementContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitSelectors is called when exiting the selectors production.
	ExitSelectors(c *SelectorsContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitUnaliasedSelector is called when exiting the unaliasedSelector production.
	ExitUnaliasedSelector(c *UnaliasedSelectorContext)

	// ExitSelectionAddition is called when exiting the selectionAddition production.
	ExitSelectionAddition(c *SelectionAdditionContext)

	// ExitSelectionMultiplication is called when exiting the selectionMultiplication production.
	ExitSelectionMultiplication(c *SelectionMultiplicationContext)

	// ExitSelectionGroup is called when exiting the selectionGroup production.
	ExitSelectionGroup(c *SelectionGroupContext)

	// ExitSelectionGroupWithField is called when exiting the selectionGroupWithField production.
	ExitSelectionGroupWithField(c *SelectionGroupWithFieldContext)

	// ExitSelectorModifier is called when exiting the selectorModifier production.
	ExitSelectorModifier(c *SelectorModifierContext)

	// ExitFieldSelectorModifier is called when exiting the fieldSelectorModifier production.
	ExitFieldSelectorModifier(c *FieldSelectorModifierContext)

	// ExitCollectionSubSelection is called when exiting the collectionSubSelection production.
	ExitCollectionSubSelection(c *CollectionSubSelectionContext)

	// ExitSelectionGroupWithoutField is called when exiting the selectionGroupWithoutField production.
	ExitSelectionGroupWithoutField(c *SelectionGroupWithoutFieldContext)

	// ExitSelectionTypeHint is called when exiting the selectionTypeHint production.
	ExitSelectionTypeHint(c *SelectionTypeHintContext)

	// ExitSelectionList is called when exiting the selectionList production.
	ExitSelectionList(c *SelectionListContext)

	// ExitSelectionMapOrSet is called when exiting the selectionMapOrSet production.
	ExitSelectionMapOrSet(c *SelectionMapOrSetContext)

	// ExitSelectionMap is called when exiting the selectionMap production.
	ExitSelectionMap(c *SelectionMapContext)

	// ExitSelectionSet is called when exiting the selectionSet production.
	ExitSelectionSet(c *SelectionSetContext)

	// ExitSelectionTupleOrNestedSelector is called when exiting the selectionTupleOrNestedSelector production.
	ExitSelectionTupleOrNestedSelector(c *SelectionTupleOrNestedSelectorContext)

	// ExitSimpleUnaliasedSelector is called when exiting the simpleUnaliasedSelector production.
	ExitSimpleUnaliasedSelector(c *SimpleUnaliasedSelectorContext)

	// ExitSelectionFunction is called when exiting the selectionFunction production.
	ExitSelectionFunction(c *SelectionFunctionContext)

	// ExitSelectionLiteral is called when exiting the selectionLiteral production.
	ExitSelectionLiteral(c *SelectionLiteralContext)

	// ExitSelectionFunctionArgs is called when exiting the selectionFunctionArgs production.
	ExitSelectionFunctionArgs(c *SelectionFunctionArgsContext)

	// ExitSident is called when exiting the sident production.
	ExitSident(c *SidentContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitRelationOrExpression is called when exiting the relationOrExpression production.
	ExitRelationOrExpression(c *RelationOrExpressionContext)

	// ExitCustomIndexExpression is called when exiting the customIndexExpression production.
	ExitCustomIndexExpression(c *CustomIndexExpressionContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitNormalInsertStatement is called when exiting the normalInsertStatement production.
	ExitNormalInsertStatement(c *NormalInsertStatementContext)

	// ExitJsonInsertStatement is called when exiting the jsonInsertStatement production.
	ExitJsonInsertStatement(c *JsonInsertStatementContext)

	// ExitJsonValue is called when exiting the jsonValue production.
	ExitJsonValue(c *JsonValueContext)

	// ExitUsingClause is called when exiting the usingClause production.
	ExitUsingClause(c *UsingClauseContext)

	// ExitUsingClauseObjective is called when exiting the usingClauseObjective production.
	ExitUsingClauseObjective(c *UsingClauseObjectiveContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitUpdateConditions is called when exiting the updateConditions production.
	ExitUpdateConditions(c *UpdateConditionsContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitDeleteSelection is called when exiting the deleteSelection production.
	ExitDeleteSelection(c *DeleteSelectionContext)

	// ExitDeleteOp is called when exiting the deleteOp production.
	ExitDeleteOp(c *DeleteOpContext)

	// ExitUsingClauseDelete is called when exiting the usingClauseDelete production.
	ExitUsingClauseDelete(c *UsingClauseDeleteContext)

	// ExitBatchStatement is called when exiting the batchStatement production.
	ExitBatchStatement(c *BatchStatementContext)

	// ExitBatchStatementObjective is called when exiting the batchStatementObjective production.
	ExitBatchStatementObjective(c *BatchStatementObjectiveContext)

	// ExitCreateAggregateStatement is called when exiting the createAggregateStatement production.
	ExitCreateAggregateStatement(c *CreateAggregateStatementContext)

	// ExitDropAggregateStatement is called when exiting the dropAggregateStatement production.
	ExitDropAggregateStatement(c *DropAggregateStatementContext)

	// ExitCreateFunctionStatement is called when exiting the createFunctionStatement production.
	ExitCreateFunctionStatement(c *CreateFunctionStatementContext)

	// ExitDropFunctionStatement is called when exiting the dropFunctionStatement production.
	ExitDropFunctionStatement(c *DropFunctionStatementContext)

	// ExitCreateKeyspaceStatement is called when exiting the createKeyspaceStatement production.
	ExitCreateKeyspaceStatement(c *CreateKeyspaceStatementContext)

	// ExitCreateTableStatement is called when exiting the createTableStatement production.
	ExitCreateTableStatement(c *CreateTableStatementContext)

	// ExitTableDefinition is called when exiting the tableDefinition production.
	ExitTableDefinition(c *TableDefinitionContext)

	// ExitTableColumns is called when exiting the tableColumns production.
	ExitTableColumns(c *TableColumnsContext)

	// ExitColumnMask is called when exiting the columnMask production.
	ExitColumnMask(c *ColumnMaskContext)

	// ExitColumnMaskArguments is called when exiting the columnMaskArguments production.
	ExitColumnMaskArguments(c *ColumnMaskArgumentsContext)

	// ExitTablePartitionKey is called when exiting the tablePartitionKey production.
	ExitTablePartitionKey(c *TablePartitionKeyContext)

	// ExitTableProperty is called when exiting the tableProperty production.
	ExitTableProperty(c *TablePropertyContext)

	// ExitTableClusteringOrder is called when exiting the tableClusteringOrder production.
	ExitTableClusteringOrder(c *TableClusteringOrderContext)

	// ExitCreateTypeStatement is called when exiting the createTypeStatement production.
	ExitCreateTypeStatement(c *CreateTypeStatementContext)

	// ExitTypeColumns is called when exiting the typeColumns production.
	ExitTypeColumns(c *TypeColumnsContext)

	// ExitCreateIndexStatement is called when exiting the createIndexStatement production.
	ExitCreateIndexStatement(c *CreateIndexStatementContext)

	// ExitIndexIdent is called when exiting the indexIdent production.
	ExitIndexIdent(c *IndexIdentContext)

	// ExitCreateMaterializedViewStatement is called when exiting the createMaterializedViewStatement production.
	ExitCreateMaterializedViewStatement(c *CreateMaterializedViewStatementContext)

	// ExitViewPrimaryKey is called when exiting the viewPrimaryKey production.
	ExitViewPrimaryKey(c *ViewPrimaryKeyContext)

	// ExitViewPartitionKey is called when exiting the viewPartitionKey production.
	ExitViewPartitionKey(c *ViewPartitionKeyContext)

	// ExitViewProperty is called when exiting the viewProperty production.
	ExitViewProperty(c *ViewPropertyContext)

	// ExitViewClusteringOrder is called when exiting the viewClusteringOrder production.
	ExitViewClusteringOrder(c *ViewClusteringOrderContext)

	// ExitCreateTriggerStatement is called when exiting the createTriggerStatement production.
	ExitCreateTriggerStatement(c *CreateTriggerStatementContext)

	// ExitDropTriggerStatement is called when exiting the dropTriggerStatement production.
	ExitDropTriggerStatement(c *DropTriggerStatementContext)

	// ExitAlterKeyspaceStatement is called when exiting the alterKeyspaceStatement production.
	ExitAlterKeyspaceStatement(c *AlterKeyspaceStatementContext)

	// ExitAlterTableStatement is called when exiting the alterTableStatement production.
	ExitAlterTableStatement(c *AlterTableStatementContext)

	// ExitIsStaticColumn is called when exiting the isStaticColumn production.
	ExitIsStaticColumn(c *IsStaticColumnContext)

	// ExitAlterMaterializedViewStatement is called when exiting the alterMaterializedViewStatement production.
	ExitAlterMaterializedViewStatement(c *AlterMaterializedViewStatementContext)

	// ExitAlterTypeStatement is called when exiting the alterTypeStatement production.
	ExitAlterTypeStatement(c *AlterTypeStatementContext)

	// ExitDropKeyspaceStatement is called when exiting the dropKeyspaceStatement production.
	ExitDropKeyspaceStatement(c *DropKeyspaceStatementContext)

	// ExitDropTableStatement is called when exiting the dropTableStatement production.
	ExitDropTableStatement(c *DropTableStatementContext)

	// ExitDropTypeStatement is called when exiting the dropTypeStatement production.
	ExitDropTypeStatement(c *DropTypeStatementContext)

	// ExitDropIndexStatement is called when exiting the dropIndexStatement production.
	ExitDropIndexStatement(c *DropIndexStatementContext)

	// ExitDropMaterializedViewStatement is called when exiting the dropMaterializedViewStatement production.
	ExitDropMaterializedViewStatement(c *DropMaterializedViewStatementContext)

	// ExitTruncateStatement is called when exiting the truncateStatement production.
	ExitTruncateStatement(c *TruncateStatementContext)

	// ExitGrantPermissionsStatement is called when exiting the grantPermissionsStatement production.
	ExitGrantPermissionsStatement(c *GrantPermissionsStatementContext)

	// ExitRevokePermissionsStatement is called when exiting the revokePermissionsStatement production.
	ExitRevokePermissionsStatement(c *RevokePermissionsStatementContext)

	// ExitGrantRoleStatement is called when exiting the grantRoleStatement production.
	ExitGrantRoleStatement(c *GrantRoleStatementContext)

	// ExitRevokeRoleStatement is called when exiting the revokeRoleStatement production.
	ExitRevokeRoleStatement(c *RevokeRoleStatementContext)

	// ExitListPermissionsStatement is called when exiting the listPermissionsStatement production.
	ExitListPermissionsStatement(c *ListPermissionsStatementContext)

	// ExitPermission is called when exiting the permission production.
	ExitPermission(c *PermissionContext)

	// ExitPermissionOrAll is called when exiting the permissionOrAll production.
	ExitPermissionOrAll(c *PermissionOrAllContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitDataResource is called when exiting the dataResource production.
	ExitDataResource(c *DataResourceContext)

	// ExitJmxResource is called when exiting the jmxResource production.
	ExitJmxResource(c *JmxResourceContext)

	// ExitRoleResource is called when exiting the roleResource production.
	ExitRoleResource(c *RoleResourceContext)

	// ExitFunctionResource is called when exiting the functionResource production.
	ExitFunctionResource(c *FunctionResourceContext)

	// ExitCreateUserStatement is called when exiting the createUserStatement production.
	ExitCreateUserStatement(c *CreateUserStatementContext)

	// ExitAlterUserStatement is called when exiting the alterUserStatement production.
	ExitAlterUserStatement(c *AlterUserStatementContext)

	// ExitDropUserStatement is called when exiting the dropUserStatement production.
	ExitDropUserStatement(c *DropUserStatementContext)

	// ExitAddIdentityStatement is called when exiting the addIdentityStatement production.
	ExitAddIdentityStatement(c *AddIdentityStatementContext)

	// ExitDropIdentityStatement is called when exiting the dropIdentityStatement production.
	ExitDropIdentityStatement(c *DropIdentityStatementContext)

	// ExitListUsersStatement is called when exiting the listUsersStatement production.
	ExitListUsersStatement(c *ListUsersStatementContext)

	// ExitCreateRoleStatement is called when exiting the createRoleStatement production.
	ExitCreateRoleStatement(c *CreateRoleStatementContext)

	// ExitAlterRoleStatement is called when exiting the alterRoleStatement production.
	ExitAlterRoleStatement(c *AlterRoleStatementContext)

	// ExitDropRoleStatement is called when exiting the dropRoleStatement production.
	ExitDropRoleStatement(c *DropRoleStatementContext)

	// ExitListRolesStatement is called when exiting the listRolesStatement production.
	ExitListRolesStatement(c *ListRolesStatementContext)

	// ExitRoleOptions is called when exiting the roleOptions production.
	ExitRoleOptions(c *RoleOptionsContext)

	// ExitRoleOption is called when exiting the roleOption production.
	ExitRoleOption(c *RoleOptionContext)

	// ExitDcPermission is called when exiting the dcPermission production.
	ExitDcPermission(c *DcPermissionContext)

	// ExitCidrPermission is called when exiting the cidrPermission production.
	ExitCidrPermission(c *CidrPermissionContext)

	// ExitUserPassword is called when exiting the userPassword production.
	ExitUserPassword(c *UserPasswordContext)

	// ExitDescribeStatement is called when exiting the describeStatement production.
	ExitDescribeStatement(c *DescribeStatementContext)

	// ExitCident is called when exiting the cident production.
	ExitCident(c *CidentContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitFident is called when exiting the fident production.
	ExitFident(c *FidentContext)

	// ExitNoncol_ident is called when exiting the noncol_ident production.
	ExitNoncol_ident(c *Noncol_identContext)

	// ExitKeyspaceName is called when exiting the keyspaceName production.
	ExitKeyspaceName(c *KeyspaceNameContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitColumnFamilyName is called when exiting the columnFamilyName production.
	ExitColumnFamilyName(c *ColumnFamilyNameContext)

	// ExitUserTypeName is called when exiting the userTypeName production.
	ExitUserTypeName(c *UserTypeNameContext)

	// ExitUserOrRoleName is called when exiting the userOrRoleName production.
	ExitUserOrRoleName(c *UserOrRoleNameContext)

	// ExitKsName is called when exiting the ksName production.
	ExitKsName(c *KsNameContext)

	// ExitCfName is called when exiting the cfName production.
	ExitCfName(c *CfNameContext)

	// ExitIdxName is called when exiting the idxName production.
	ExitIdxName(c *IdxNameContext)

	// ExitRoleName is called when exiting the roleName production.
	ExitRoleName(c *RoleNameContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitFullMapLiteral is called when exiting the fullMapLiteral production.
	ExitFullMapLiteral(c *FullMapLiteralContext)

	// ExitSetOrMapLiteral is called when exiting the setOrMapLiteral production.
	ExitSetOrMapLiteral(c *SetOrMapLiteralContext)

	// ExitSetLiteral is called when exiting the setLiteral production.
	ExitSetLiteral(c *SetLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitCollectionLiteral is called when exiting the collectionLiteral production.
	ExitCollectionLiteral(c *CollectionLiteralContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitUsertypeLiteral is called when exiting the usertypeLiteral production.
	ExitUsertypeLiteral(c *UsertypeLiteralContext)

	// ExitTupleLiteral is called when exiting the tupleLiteral production.
	ExitTupleLiteral(c *TupleLiteralContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitIntValue is called when exiting the intValue production.
	ExitIntValue(c *IntValueContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitAllowedFunctionName is called when exiting the allowedFunctionName production.
	ExitAllowedFunctionName(c *AllowedFunctionNameContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFunctionArgs is called when exiting the functionArgs production.
	ExitFunctionArgs(c *FunctionArgsContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitTermAddition is called when exiting the termAddition production.
	ExitTermAddition(c *TermAdditionContext)

	// ExitTermMultiplication is called when exiting the termMultiplication production.
	ExitTermMultiplication(c *TermMultiplicationContext)

	// ExitTermGroup is called when exiting the termGroup production.
	ExitTermGroup(c *TermGroupContext)

	// ExitSimpleTerm is called when exiting the simpleTerm production.
	ExitSimpleTerm(c *SimpleTermContext)

	// ExitColumnOperation is called when exiting the columnOperation production.
	ExitColumnOperation(c *ColumnOperationContext)

	// ExitColumnOperationDifferentiator is called when exiting the columnOperationDifferentiator production.
	ExitColumnOperationDifferentiator(c *ColumnOperationDifferentiatorContext)

	// ExitNormalColumnOperation is called when exiting the normalColumnOperation production.
	ExitNormalColumnOperation(c *NormalColumnOperationContext)

	// ExitShorthandColumnOperation is called when exiting the shorthandColumnOperation production.
	ExitShorthandColumnOperation(c *ShorthandColumnOperationContext)

	// ExitCollectionColumnOperation is called when exiting the collectionColumnOperation production.
	ExitCollectionColumnOperation(c *CollectionColumnOperationContext)

	// ExitUdtColumnOperation is called when exiting the udtColumnOperation production.
	ExitUdtColumnOperation(c *UdtColumnOperationContext)

	// ExitColumnCondition is called when exiting the columnCondition production.
	ExitColumnCondition(c *ColumnConditionContext)

	// ExitProperties is called when exiting the properties production.
	ExitProperties(c *PropertiesContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitPropertyValue is called when exiting the propertyValue production.
	ExitPropertyValue(c *PropertyValueContext)

	// ExitRelationType is called when exiting the relationType production.
	ExitRelationType(c *RelationTypeContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitContainsOperator is called when exiting the containsOperator production.
	ExitContainsOperator(c *ContainsOperatorContext)

	// ExitInMarker is called when exiting the inMarker production.
	ExitInMarker(c *InMarkerContext)

	// ExitTupleOfIdentifiers is called when exiting the tupleOfIdentifiers production.
	ExitTupleOfIdentifiers(c *TupleOfIdentifiersContext)

	// ExitSingleColumnInValues is called when exiting the singleColumnInValues production.
	ExitSingleColumnInValues(c *SingleColumnInValuesContext)

	// ExitTupleOfTupleLiterals is called when exiting the tupleOfTupleLiterals production.
	ExitTupleOfTupleLiterals(c *TupleOfTupleLiteralsContext)

	// ExitMarkerForTuple is called when exiting the markerForTuple production.
	ExitMarkerForTuple(c *MarkerForTupleContext)

	// ExitTupleOfMarkersForTuples is called when exiting the tupleOfMarkersForTuples production.
	ExitTupleOfMarkersForTuples(c *TupleOfMarkersForTuplesContext)

	// ExitInMarkerForTuple is called when exiting the inMarkerForTuple production.
	ExitInMarkerForTuple(c *InMarkerForTupleContext)

	// ExitComparatorType is called when exiting the comparatorType production.
	ExitComparatorType(c *ComparatorTypeContext)

	// ExitNative_type is called when exiting the native_type production.
	ExitNative_type(c *Native_typeContext)

	// ExitCollection_type is called when exiting the collection_type production.
	ExitCollection_type(c *Collection_typeContext)

	// ExitTuple_type is called when exiting the tuple_type production.
	ExitTuple_type(c *Tuple_typeContext)

	// ExitVector_type is called when exiting the vector_type production.
	ExitVector_type(c *Vector_typeContext)

	// ExitUsername is called when exiting the username production.
	ExitUsername(c *UsernameContext)

	// ExitIdentity is called when exiting the identity production.
	ExitIdentity(c *IdentityContext)

	// ExitMbean is called when exiting the mbean production.
	ExitMbean(c *MbeanContext)

	// ExitNon_type_ident is called when exiting the non_type_ident production.
	ExitNon_type_ident(c *Non_type_identContext)

	// ExitUnreserved_keyword is called when exiting the unreserved_keyword production.
	ExitUnreserved_keyword(c *Unreserved_keywordContext)

	// ExitUnreserved_function_keyword is called when exiting the unreserved_function_keyword production.
	ExitUnreserved_function_keyword(c *Unreserved_function_keywordContext)

	// ExitBasic_unreserved_keyword is called when exiting the basic_unreserved_keyword production.
	ExitBasic_unreserved_keyword(c *Basic_unreserved_keywordContext)
}
