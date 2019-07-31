// The env package is a thin wrapper around the os package environment variable functions. The main use case is to determine the "environment type". I.e., 'dev', 'test', or 'production'. env also provides 'Must' flavor which panic if the environment variable is not set. Finally, we go ahead and wrap the basic set and unset functions (with proper camel casing!) while we're at it.
package env
