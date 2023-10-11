package latex

import (
	"encoding/json"
	"os"
)

var (
	// PreDefined holds set options will be loaded on init, but can be
	// redone later
	GlobalPreDefined *PreDef
	// default path
	predDefDefPath = "default_env.json"
)

func init() {
	LoadPreDefined(predDefDefPath)
}

/*
OptionHolder is a parentclass to easy provide a wrapping of strings with
Options
*/
type OptionHolder struct {
	// Options see Option struct
	Options []*Option
	// Holds Option by name or type to be looked up
	OptionsByName
}

/*
WrapWithOptions apply with options
*/
func (oh *OptionHolder) WrapWithOptions(text string) string {
	// add options for name and type if any found
	for _, opt := range oh.GetInnerOptions() {
		text = opt.WrapString(text)
	}
	for _, opt := range oh.Options {
		text = opt.WrapString(text)
	}
	return text
}

/*
AddOption adds a option with name and optional parameter
*/
func (oh *OptionHolder) AddOption(
	env string, opts, params []string, short bool, pre string, inv bool) {
	oh.Options = append(oh.Options, &Option{env, opts, params, short, pre, inv})
}

/*
Option holds a option for a LatexStr
*/
type Option struct {
	// Env options name/environment
	Env string `json:"env"`
	// Opts holds additional parameter
	Opts []string `json:"opts"`
	// Args holds the arguments for this environment
	Args []string `json:"args"`
	// UseShort hanges wrapper mode to short mode with Brackets. Params are
	// not used in short mode
	UseShort bool `json:"use_short"`
	// PreDefined holds a string, if it matches a predefined, that one is used
	Pre string `json:"pre"`
	// Invert is used to suspend option, use with care!
	Invert bool `json:"invert"`
}

/*
Wraps a string with given option
*/
func (o *Option) WrapString(text string) string {
	if GlobalPreDefined != nil && GlobalPreDefined.PreDefs != nil {
		if val, ok := GlobalPreDefined.PreDefs[o.Pre]; ok {
			return val.WrapString(text)
		}
	}

	// use short
	if o.UseShort {
		ret := LC + o.Env
		for _, arg := range o.Args {
			ret += Brackets(arg)
		}
		ret += Brackets(text)
		return ret
	}

	// begin keyword
	begin := Begin + Brackets(o.Env) + JoinParams(o.Opts...)
	for _, arg := range o.Args {
		begin += Brackets(arg)
	}
	begin += LineBr

	text = CheckAddLineBreak(text)

	// end keyword
	end := End + Brackets(o.Env) + LineBr

	// inverted, so end it for block and begin it to continue
	if o.Invert {
		return end + text + begin
	}

	// return value text with \begin \end block
	return begin + text + end
}

// #############################################################################
// #							Options by Name
// #############################################################################

/*
Holds default Options that are added to parts of the late document.
*/
type OptionsByName struct {
	// Name (Headline, etc.)
	Name string
	// Type
	Type string
	// Lookup for parent nodes to find find first none empty map
	Parent *OptionsByName
	// InnerOpts for block options defined by either name or type
	// (e.g. section)
	InnerOpts map[string][]*Option `json:"inner_opts"`
	// OuterOpts wraps the whole node in option (name and type)
	OuterOpts map[string][]*Option `json:"outer_opts"`
	// HeadLineOptions are also lookedup by nme and type
	HeadLineOptions map[string][]*Option `json:"headline_opts"`
}

const (
	inner  = 0
	outer  = 1
	hlopts = 2
)

/*
AddDefaultOptions returns the configurable default options for type and name
*/
func (on *OptionsByName) GetInnerOptions() []*Option {
	return on.lookUPOpts(on.Type, on.Name, inner)
}

/*
AddDefaultOptions returns the configurable default options for type and name
*/
func (on *OptionsByName) GetOuterOptions() []*Option {
	return on.lookUPOpts(on.Type, on.Name, outer)
}

/*
GetHeadlineOption returns options to be used for headline
*/
func (on *OptionsByName) GetHeadlineOptions() []*Option {
	return on.lookUPOpts(on.Type, on.Name, hlopts)
}

func (on *OptionsByName) lookUPOpts(typ, name string, whichOne int) []*Option {

	lMap := on.InnerOpts
	if whichOne == outer {
		lMap = on.OuterOpts
	}
	if whichOne == hlopts {
		lMap = on.HeadLineOptions
	}

	ret := []*Option{}
	if val, ok := lMap[name]; ok {
		ret = append(ret, val...)
	}
	if val, ok := lMap[typ]; ok {
		ret = append(ret, val...)
	}
	if on.Parent != nil {
		ret = append(ret, on.Parent.lookUPOpts(typ, name, whichOne)...)
	}
	return ret
}

/*
SetNameTypeAndParent sets name, type and parent of a OptionHolder via
OptionsByName. An option tree is set automatically with Doc as root as long as
`Add` Method are used to create the document
*/
func (on *OptionsByName) SetNameTypeAndParent(
	name, typ string, parent *OptionsByName) {

	on.Name = name
	on.Type = typ
	on.Parent = parent
}

/*
GetDefaultsFromFile loads a json file and fills it parsed options into
*/
func (on *OptionsByName) GetDefaultsFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &on); err != nil {
		return err
	}
	return nil
}

// #############################################################################
// #							Predefined
// #############################################################################

type PreDef struct {
	PreDefs map[string]*Option `json:"predefined"`
}

/*
LoadPreDefined loads a json file and sets global pre defined short opts
*/
func LoadPreDefined(path string) error {
	data, err := os.ReadFile(path)
	GlobalPreDefined = &PreDef{}
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, GlobalPreDefined); err != nil {
		return err
	}

	return nil
}

// #############################################################################
// #							Util/Shorts
// #############################################################################

func (oh *OptionHolder) ShortOpt(name string) {
	oh.AddOption("", nil, nil, false, name, false)
}
