package log

/*
	This go file centralizes log messages (in different levels) so we have them all in one place.
	Although having the names of the consts as the error code (i.e CAKC001E) and not as a descriptive name (i.e WriteAccessTokenError)
	can reduce readability of the code that raises the error, we decided to do so for the following reasons:
		1.  Improves supportability – when we get this code in the log we can find it directly in the code without going
			through the “log_messages.go” file first
		2. Validates we don’t have error code duplications – If the code is only in the string then 2 errors can have the
			same code (which is something that a developer can easily miss). However, If they are in the object name
			then the compiler will not allow it.
*/

// ERROR MESSAGES
const CAIC001E string = "CAIC001E Error getting the AWS resource type from environment variable '%s' or flag '-%s'"
const CAIC002E string = "CAIC002E Error getting the AWS resource type. %s"
const CAIC003E string = "CAIC003E Failed to retrieve the AWS IAM Credential for resource '%s'. %s"
const CAIC004E string = "CAIC004E Failed to get Conjur authentication request. %s"
const CAIC005E string = "CAIC005E Failed to initialize Conjur HTTP Client. %s"
const CAIC006E string = "CAIC006E Error getting Conjur authentication url environment variable '%s'  or flag '-%s'"
const CAIC007E string = "CAIC007E Failed to authenticate to Conjur. %s"
const CAIC008E string = "CAIC008E Failure to load Conjur config. %s"
const CAIC009E string = "CAIC009E Error getting the Conjur login environment variable '%s' or flag '-%s'"
const CAIC010E string = "CAIC010E Failed to write access token to '%s'. %s"
const CAIC011E string = "CAIC011E Failed to retrieve secret '%s'. %s"
const CAIC012E string = "CAIC012E Error getting the Conjur account environment variables '%s' or flag '-%s'"
const CAIC013E string = "CAIC012E Error getting the Conjur appliance URL environment variables '%s' or flag '-%s'"

// INFO MESSAGES
const CAIC001I string = "CAIC001I Retrieving AWS IAM credential for resource type '%s'"
const CAIC002I string = "CAIC002I Successfully retrieved AWS IAM credential"
const CAIC003I string = "CAIC003I Successfully created the Conjur authentication request"
const CAIC004I string = "CAIC004I Attempting to authenticate to conjur: authnUrl=%s, login=%s, account=%s"
const CAIC005I string = "CAIC005I Attempting to authenticate to conjur '%s'"
const CAIC006I string = "CAIC006I Successfully authenticated to conjur with login '%s'"
const CAIC007I string = "CAIC007I Successfully wrote access token to '%s'"

// DEBUG MESSAGES
const CAIC001D string = "CAIC001D Debug mode is enabled"
