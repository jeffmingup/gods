// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import "github.com/emirpasic/gods/containers"

func assertEnumerableImplementation() {
	var _ containers.EnumerableWithIndex = (*List)(nil)
}

