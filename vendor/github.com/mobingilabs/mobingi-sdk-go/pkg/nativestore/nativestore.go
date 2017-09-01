package nativestore

import dcred "github.com/docker/docker-credential-helpers/credentials"

func Set(lbl, url, user, secret string) error {
	cr := &dcred.Credentials{
		ServerURL: url,
		Username:  user,
		Secret:    secret,
	}

	dcred.SetCredsLabel(lbl)
	return ns.Add(cr)
}

func Get(url string) (string, string, error) {
	return ns.Get(url)
}
