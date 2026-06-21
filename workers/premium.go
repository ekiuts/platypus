package workers

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/farghul/platypus/internal"
)

// Gather the premium plugin semantic versions currently installed and available
func subscription() string {
	var spotlightInstalled = current("freemius/spotlight-social-photo-feeds-premium")
	var exportInstalled = current("premium-plugin/wp-all-export-pro")
	var polylangInstalled = current("premium-plugin/polylang-pro")
	var spotlightAvailable = latest(internal.Changelog.Spotlight, "<h2>v")
	var exportAvailable = latest(internal.Changelog.WPExport, "<h4>")
	var polylangAvailable = latest(internal.Changelog.Poly, "<pre")
	collect := results(spotlightInstalled, spotlightAvailable, "spotlight-social-photo-feeds-premium") + results(exportInstalled, exportAvailable, "wp-all-export-pro") + results(polylangInstalled, polylangAvailable, "polylang-pro")
	return collect
}

// Determine if a WordPress core update is needed
func wpcore() string {
	var coreInstalled = current("roots/wordpress")
	var coreAvailable = sift()
	collect := results(coreInstalled, coreAvailable, "wordpress")
	return collect
}

// Print the results if an update is available
func results(current, update, plugin string) string {
	var status string

	if comparison(current, update) < 0 {
		if plugin == "wordpress" {
			status = "roots/" + plugin + ":" + update + "\n"
		} else {
			status = "premium-plugin/" + plugin + ":" + update + "\n"
		}
	}
	return status
}

// Find the current semantic version of our premium plugins from the composer.json file
func current(p string) string {
	where := strings.TrimSuffix(internal.Ecosystem.Install, "web/wp") + "composer.json"
	what := internal.Concat("ssh", "-T", internal.Ecosystem.User+"@"+internal.Ecosystem.Server, " cat "+where)
	internal.Inspect(os.WriteFile(internal.Remains[1], what, 0666))
	grep, _ := internal.Capture("grep", p, internal.Remains[1])
	return regmatch(strings.TrimSpace(string(grep)))
}

// Find the latest semantic version of our premium plugins from the applicable websites
func latest(u, g string) string {
	internal.Capture("curl", "-s", u, "-o", internal.Remains[2])
	grep, _ := internal.Capture("grep", g, internal.Remains[2])
	internal.Inspect(os.WriteFile(internal.Remains[0], grep, 0666))
	head, _ := internal.Capture("head", "-n 1", internal.Remains[0])
	return regmatch(strings.TrimSpace(string(head)))
}

// Directly compare semantic versioning
func comparison(v1, v2 string) int {
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	maxLen := max(len(parts2), len(parts1))

	for i := range maxLen {
		var p1, p2 int

		if i < len(parts1) {
			p1, _ = strconv.Atoi(parts1[i])
		}

		if i < len(parts2) {
			p2, _ = strconv.Atoi(parts2[i])
		}

		if p1 > p2 {
			return 1
		} else if p1 < p2 {
			return -1
		}
	}

	return 0
}

// Remove all extraneous material, leaving only the version number itself
func regmatch(p string) string {
	var match []string
	tri := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}`)
	quad := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}.\d{1,3}`)

	if quad.MatchString(p) {
		match = quad.FindAllString(p, -1)
	} else if tri.MatchString(p) {
		match = tri.FindAllString(p, -1)
	}
	return match[0]
}
