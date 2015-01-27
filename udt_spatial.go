package mssql

import ()

// TODO: translate from https://github.com/patriksimek/node-mssql/blob/master/src/udt.coffee
func decodeGeography(data []byte) interface{} {
	return "geography"
}

func decodeGeometry(data []byte) interface{} {
	return "geometry"
}
