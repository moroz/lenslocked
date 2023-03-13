package main

import "github.com/moroz/lenslocked/utils"

var DATABASE_URL string = utils.RequireEnvVar("DATABASE_URL")
var CSRF_SECRET []byte = utils.RequireHexEnvVar("CSRF_SECRET")

const COOKIE_KEY = "_lenslocked_session"
