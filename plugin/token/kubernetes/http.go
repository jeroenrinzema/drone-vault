// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func post(path string, in, out interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		return err
	}

	transporter := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transporter}
	res, err := client.Post(path, "application/json", buf)
	fmt.Println(in)
	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return errors.New(
			res.Status,
		)
	}
	return json.NewDecoder(res.Body).Decode(out)
}
