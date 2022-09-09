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
	// https://buildnumbers.wordpress.com/sharepoint/

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

	// SharePoint Server 2016
	s.Builds["16.0.5032"] = Build{"July 2020 CU", "2020 July 14", "SharePoint Server 2016 (language independent)", "KB4484436, KB4484440"}
	s.Builds["16.0.5017"] = Build{"July 2020 CU", "2020 July 14", "SharePoint Server 2016 (language independent)", "KB4484402, KB4484344"}
	s.Builds["16.0.5005"] = Build{"May 2020 CU", "2020 May 12", "SharePoint Server 2016 (language independent)", "KB4484336"}
	s.Builds["16.0.4993"] = Build{"April 2020 CU", "2020 April 14", "SharePoint Server 2016 (language independent)", "KB4484299, KB4484301"}
	s.Builds["16.0.4978"] = Build{"March 2020 CU", "2020 March 10", "SharePoint Server 2016 (language independent)", "KB4484272, KB4484275"}
	s.Builds["16.0.4966"] = Build{"February 2020 CU", "2020 February 11", "SharePoint Server 2016 (language independent)", "KB4484255, KB4484257"}
	s.Builds["16.0.4954"] = Build{"January 2020 CU", "2020 January 14", "SharePoint Server 2016 (language independent)", "KB4484215, KB4484220"}
	s.Builds["16.0.4939"] = Build{"December 2019 CU", "2019 December 10", "SharePoint Server 2016 (language independent)", "KB4484178, KB4484181"}
	s.Builds["16.0.4927"] = Build{"November 2019 CU", "2019 November 12", "SharePoint Server 2016 (language independent)", "KB4484143, KB4484147"}
	s.Builds["16.0.4912"] = Build{"October 2019 CU", "2019 October 8", "SharePoint Server 2016 (language independent)", "KB4484111, KB4484115"}
	s.Builds["16.0.4900"] = Build{"September 2019 CU", "2019 September 10", "SharePoint Server 2016 (language independent)", "KB4475590, KB4475594"}
	s.Builds["16.0.4888"] = Build{"August 2019 CU", "2019 August 13", "SharePoint Server 2016 (language independent)", "KB4475549, KB4464553"}
	s.Builds["16.0.4873"] = Build{"July 2019 CU", "2019 July 9", "SharePoint Server 2016 (language independent)", "KB4475520"}
	s.Builds["16.0.4861"] = Build{"June 2019 CU", "2019 June 11", "SharePoint Server 2016 (language independent)", "KB4464594"}
	s.Builds["16.0.4849"] = Build{"May 2019 CU", "2019 May 14", "SharePoint Server 2016 (language independent)", "KB4464549"}
	s.Builds["16.0.4834"] = Build{"April 2019 CU", "2019 April 9", "SharePoint Server 2016 (language independent)", "KB4464510, KB4461507"}
	s.Builds["16.0.4822"] = Build{"March 2019 CU", "2019 March 12", "SharePoint Server 2016 (language independent)", "KB4462211"}
	s.Builds["16.0.4810"] = Build{"February 2019 CU", "2019 February 8", "SharePoint Server 2016 (language independent)", "KB4462155"}
	s.Builds["16.0.4795"] = Build{"January 2019 CU", "2019 January 8", "SharePoint Server 2016 (language independent)", "KB4461598"}
	s.Builds["16.0.4783"] = Build{"December 2018 CU", "2018 December 11", "SharePoint Server 2016 (language independent)", "KB4461541"}
	s.Builds["16.0.4771"] = Build{"November 2018 CU", "2018 November 13", "SharePoint Server 2016 (language independent)", "KB4461501"}
	s.Builds["16.0.4756"] = Build{"October 2018 CU", "2018 October 9", "SharePoint Server 2016 (language independent)", "KB4461447,KB4092463"}
	s.Builds["16.0.4744"] = Build{"September 2018 CU", "2018 September 11", "SharePoint Server 2016 (language independent)", "KB4092459"}
	s.Builds["16.0.4732"] = Build{"August 2018 CU", "2018 August 14", "SharePoint Server 2016 (language independent)", "KB4032256,KB4022231"}
	s.Builds["16.0.4717"] = Build{"July 2018 CU", "2018 July 10", "SharePoint Server 2016 (language independent)", "KB4022228"}
	s.Builds["16.0.4705"] = Build{"June 2018 CU", "2018 June 12", "SharePoint Server 2016 (language independent)", "KB4022173,KB4022178"}
	s.Builds["16.0.4690"] = Build{"May 2018 CU", "2018 May 8", "SharePoint Server 2016 (language independent)", "KB4018381,KB4018386"}
	s.Builds["16.0.4678"] = Build{"April 2018 CU", "2018 April 10", "SharePoint Server 2016 (language independent)", "KB4018336,KB4018340"}
	s.Builds["16.0.4666"] = Build{"March 2018 CU", "2018 March 13", "SharePoint Server 2016 (language independent)", "KB4018293,KB4011687"}
	s.Builds["16.0.4654"] = Build{"February 2018 CU", "2018 February 13", "SharePoint Server 2016 (language independent)", "KB4011680"}
	s.Builds["16.0.4639"] = Build{"January 2018 CU", "2018 January 9", "SharePoint Server 2016 (language independent)", "KB4011642,KB4011645"}
	s.Builds["16.0.4627"] = Build{"December 2017 CU", "2017 December 12", "SharePoint Server 2016 (language independent)", "KB4011576,KB4011578"}
	s.Builds["16.0.4615"] = Build{"November 2017 CU", "2017 November 14", "SharePoint Server 2016 (language independent)", "KB4011244,KB4011243"}
	s.Builds["16.0.4600"] = Build{"October 2017 CU", "2017 October 10", "SharePoint Server 2016 (language independent)", "KB4011217,KB4011161"}
	s.Builds["16.0.4588"] = Build{"September 2017 CU aka Feature Pack 2", "2017 September 12", "SharePoint Server 2016 (language independent)", "KB4011127,KB4011112"}
	s.Builds["16.0.4573"] = Build{"August 2017 CU", "2017 August 8", "SharePoint Server 2016 (language independent)", "KB4011049,KB4011053"}
	s.Builds["16.0.4561"] = Build{"July 2017 CU", "2017 July 11", "SharePoint Server 2016", "KB3213544,KB3213543"}
	s.Builds["16.0.4549"] = Build{"June 2017 CU", "2017 June 13", "SharePoint Server 2016 (language independent)", "KB3203432,KB3203433"}
	s.Builds["16.0.4534"] = Build{"May 2017 CU", "2017 May 9", "SharePoint Server 2016 (language independent)", "KB3191880,KB3191884"}
	s.Builds["16.0.4522"] = Build{"April 2017 CU", "2017 April 11", "SharePoint Server 2016 (language independent)", "KB3178718,KB3178721"}
	s.Builds["16.0.4522"] = Build{"April 2017 CU", "2017 April 11", "SharePoint Server 2016 (language independent)", "KB3178718,KB3178721"}

}
