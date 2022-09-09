package main

import "fmt"

type Build struct {
	Name        string
	Date        string
	Component   string
	Description string
}

func (b Build) Print() {
	fmt.Println("\tBuild Name: " + b.Name)
	fmt.Println("\tRelease Date: " + b.Date)
	fmt.Println("\tComponent: " + b.Component)
	fmt.Println("\tDescription: " + b.Description)
}

func (s *SharePointScanner) PopulateBuilds() {
	s.Builds = make(map[string]Build)

	// SharePoint Server 2019
	s.Builds["16.0.10363"] = Build{"July 2020 CU", "2020 July 14", "SharePoint Server 2019 (language independent)", "KB4484453"}
	s.Builds["16.0.10362"] = Build{"July 2020 CU", "2020 July 14", "SharePoint Server 2019 (language independent)", "KB4484453"}
	s.Builds["16.0.10361"] = Build{"June 2020 CU", "2020 June 9", "SharePoint Server 2019 (language independent)", "KB4484400"}
	s.Builds["16.0.10360"] = Build{"June 2020 CU", "2020 June 9", "SharePoint Server 2019 (language independent)", "KB4484404"}
	s.Builds["16.0.10359"] = Build{"May 2020 CU", "May 2020 CU", "SharePoint Server 2019 (language independent)", "KB4484332, KB4484331"}
	s.Builds["16.0.10358"] = Build{"April 2020 CU", "2020 April 14", "SharePoint Server 2019 (language independent)", "KB4484332, KB4484291"}
	s.Builds["16.0.10357"] = Build{"March 2020 CU", "2020 March 10", "SharePoint Server 2019 (language independent)", "KB4484271, KB4484277"}
	s.Builds["16.0.10355"] = Build{"February 2020 CU", "2020 February 11", "SharePoint Server 2019 (language independent)", "KB4484259, KB4484225"}
	s.Builds["16.0.10354"] = Build{"January 2020 CU", "2020 January 14", "SharePoint Server 2019 (language independent)", "KB4484224"}
	s.Builds["16.0.10353"] = Build{"December 2019 CU", "2019 December 10", "SharePoint Server 2019 (language independent)", "KB4484177, KB4484176"}
	s.Builds["16.0.10352"] = Build{"November 2019 CU", "2019 November 12", "SharePoint Server 2019 (language independent)", "KB4484142, KB4484149"}
	s.Builds["16.0.10351"] = Build{"October 2019 CU", "2019 October 8", "SharePoint Server 2019 (language independent)", "KB4484110, KB4484109"}
	s.Builds["16.0.10350"] = Build{"September 2019 CU", "2019 September 10", "SharePoint Server 2019 (language independent)", "KB4475596, KB4464557"}
	s.Builds["16.0.10349"] = Build{"August 2019 CU", "2019 August 13", "SharePoint Server 2019 (language independent)", "KB4475555"}
	s.Builds["16.0.10348"] = Build{"July 2019 CU", "2019 July 9", "SharePoint Server 2019 (language independent)", "KB4475529"}
	s.Builds["16.0.10346"] = Build{"June 2019 CU", "2019 June 11", "SharePoint Server 2019 (language independent)", "KB4475512"}
	s.Builds["16.0.10345"] = Build{"May 2019 CU", "2019 May 14", "SharePoint Server 2019 (language independent)", "KB4464556"}
	s.Builds["16.0.10343"] = Build{"April 2019 CU", "2019 April 9", "SharePoint Server 2019 (language independent)", "KB4464518,KB4462221"}
	s.Builds["16.0.10342"] = Build{"March 2019 CU", "2019 March 12", "SharePoint Server 2019 (language independent)", "KB4462199"}
	s.Builds["16.0.10341"] = Build{"February 2019 CU", "2019 February 8", "SharePoint Server 2019 (language independent)", "KB4462171,KB4462170"}
	s.Builds["16.0.10340"] = Build{"January 2019 CU", "2019 January 8", "SharePoint Server 2019 (language independent)", "KB4461634,KB4461514"}
	s.Builds["16.0.10339"] = Build{"December 2018 CU", "2018 December 11", "SharePoint Server 2019 (language independent)", "KB4461548,KB4461548"}
	s.Builds["16.0.10338"] = Build{"RTM", "2018 October 22", "SharePoint Server 2019 (language independent)", "KB4461513"}
	s.Builds["16.0.10337"] = Build{"Public Preview", "2018 July 24", "SharePoint Server 2019", "N/A"}
	s.Builds["16.0.10711"] = Build{"Public Preview", "2018 July 24", "SharePoint Server 2019", "N/A"}

}
