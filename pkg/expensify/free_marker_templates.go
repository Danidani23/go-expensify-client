package expensify

import (
	"fmt"
	"strings"
)

func generateFreeMarkerTemplate(reportFieldNames []string, expenseFieldNames []string) string {
	// Start building the FreeMarker template
	var sb strings.Builder

	// start of the report part
	sb.WriteString("[")
	sb.WriteString("<#list reports as report>\n")
	sb.WriteString("    <#-- Here, data at the report level is accessed using the \"report\" object -->\n")
	sb.WriteString("{\n")

	// Loop through each report field ID and add it to the template
	for i, fieldID := range reportFieldNames {
		switch fieldID {
		case "actionList":
			sb.WriteString(fmt.Sprintf("    \"%s\": [<#list report.%s as action>{\"action\": \"${action.action}\"}<#if action_has_next>,</#if></#list>]", fieldID, fieldID))
		case "approvers":
			sb.WriteString(fmt.Sprintf("    \"%s\": [<#list report.%s as approver>{\"email\": \"${approver.email}\", \"name\": \"${approver.fullName}\"}<#if approver_has_next>,</#if></#list>]", fieldID, fieldID))
		case "total":
			sb.WriteString(fmt.Sprintf("    \"%s\": ${report.%s / 100}", fieldID, fieldID))
		case "isACHReimbursed":
			sb.WriteString(fmt.Sprintf("    \"%s\": ${report.%s?string(\"true\", \"false\")}", fieldID, fieldID))
		case "transactionList":
			// start of the expense part
			sb.WriteString("    \"transactionList\":[<#list report.transactionList as expense>\n")
			sb.WriteString("        {\n")

			if len(expenseFieldNames) > 0 {
				for i, expenseFieldID := range expenseFieldNames {
					sb.WriteString(fmt.Sprintf("        \"%s\": \"${expense.%s}\"", expenseFieldID, expenseFieldID))
					if i < len(expenseFieldNames)-1 {
						sb.WriteString(",\n")
					} else {
						sb.WriteString("\n")
					}

				}
			}

			// closing of the expense part
			sb.WriteString("        }<#if expense_has_next>,</#if>\n")
			sb.WriteString("    </#list>]\n")

		default:
			sb.WriteString(fmt.Sprintf("    \"%s\": \"${report.%s?if_exists?json_string}\"", fieldID, fieldID))
		}

		// Add a comma after each report field, except the last one
		if i < len(reportFieldNames)-1 {
			sb.WriteString(",\n")
		} else {
			sb.WriteString("\n")
		}
	}

	// closing of the report part
	sb.WriteString("}<#if report_has_next>,</#if>\n")
	sb.WriteString("</#list>\n")
	sb.WriteString("]")

	return sb.String()
}
