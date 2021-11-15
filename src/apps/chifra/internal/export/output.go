package exportPkg

/*-------------------------------------------------------------------------------------------
 * qblocks - fast, easily-accessible, fully-decentralized data from blockchains
 * copyright (c) 2016, 2021 TrueBlocks, LLC (http://trueblocks.io)
 *
 * This program is free software: you may redistribute it and/or modify it under the terms
 * of the GNU General Public License as published by the Free Software Foundation, either
 * version 3 of the License, or (at your option) any later version. This program is
 * distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details. You should have received a copy of the GNU General
 * Public License along with this program. If not, see http://www.gnu.org/licenses/.
 *-------------------------------------------------------------------------------------------*/
/*
 * Parts of this file were auto generated with makeClass --gocmds. Be careful when editing
 * it to only edit those parts of the code inside of // EXISTING_CODE blocks
 */

import (
	"net/http"

	"github.com/spf13/cobra"
)

var Options ExportOptions

func RunExport(cmd *cobra.Command, args []string) error {
	Options.Addrs = args // also Topics and FourBytes
	opts := Options

	err := opts.ValidateExport()
	if err != nil {
		return err
	}

	return opts.Globals.PassItOn("acctExport", opts.ToDashStr())
}

func ServeExport(w http.ResponseWriter, r *http.Request) {
	opts := FromRequest(w, r)

	err := opts.ValidateExport()
	if err != nil {
		opts.Globals.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	opts.Globals.PassItOn("acctExport", opts.ToDashStr())
}