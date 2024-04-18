package traffic

type Measurements struct {
	num_of_subs        float64
	traffic_per_subs_1 float64
	traffic_1          float64

	traffic_2 float64
	bhca      float64
	mht       float64

	subs_vlr           float64
	traffic_per_subs_2 float64
	switch_traffic     float64
}

type Calculations struct {
	mesurement Measurements
}

/*===========================================================*/
/* 						Implementation						 */
/*===========================================================*/

/*
	Formula: num_of_subs = traffic / (traffic_per_subs / 1000)
*/

// Number-of-Subs vs Traffic-Per-Subs vs Traffic
func (c Calculations) Calc_num_of_subs(traff float64, traff_per_subs float64) float64 {
	c.mesurement.num_of_subs = traff / (traff_per_subs / 1000)
	return Round(c.mesurement.num_of_subs, 3)
}

// Traffic vs Number-of-Subs vs Traffic-Per-Subs
func (c Calculations) Calc_traffic(no_of_subs float64, traff_per_subs float64) float64 {
	c.mesurement.traffic_1 = no_of_subs * (traff_per_subs / 1000)
	return Round(c.mesurement.traffic_1, 3)
}

// Number-of-Subs vs Traffic vs Traffic-Per-Subs
func (c Calculations) Calc_traffic_per_subs(traff float64, traff_per_subs uint) float64 {
	c.mesurement.traffic_per_subs_1 = traff * 1000 / float64(traff_per_subs)
	return Round(c.mesurement.traffic_per_subs_1, 3)
}

/*
	Formula: traffic = BHCA * MHT / 3600
*/

// traffic vs BHCA vs MHT
func (c Calculations) Calc_traffic_2(bhca float64, mht float64) float64 {
	c.mesurement.traffic_2 = (bhca * mht) / 3600
	return Round(c.mesurement.traffic_2, 3)
}

func (c Calculations) Calc_BHCA(traff float64, mht float64) float64 {
	c.mesurement.bhca = (traff * 3600) / mht
	return Round(c.mesurement.bhca, 3)
}

func (c Calculations) Calc_MHT(traff float64, bhca float64) float64 {
	c.mesurement.mht = (traff * 3600) / bhca
	return Round(c.mesurement.mht, 3)
}

/*
	Formula: subs in VLR = traffic per subs * switch traffic
*/

// Subs VLR vs traffic/subs vs switch traffic
func (c Calculations) Calc_subs_vlr(traffic_per_subs_2 float64, switch_traffic float64) float64 {
	c.mesurement.subs_vlr = traffic_per_subs_2 * switch_traffic
	return Round(c.mesurement.subs_vlr, 3)
}

// traffic/subs vs subs vlr vs switch traffic
func (c Calculations) Calc_traffic_per_subs_2(subs_vlr float64, switch_traffic float64) float64 {
	c.mesurement.traffic_per_subs_2 = subs_vlr / switch_traffic
	return Round(c.mesurement.traffic_per_subs_2, 3)
}

// switch traffic vs traffic/subs vs subs vlr
func (c Calculations) Calc_switch_traffic(subs_vlr float64, traffic_per_subs_2 float64) float64 {
	c.mesurement.switch_traffic = subs_vlr / traffic_per_subs_2
	return Round(c.mesurement.switch_traffic, 3)
}
