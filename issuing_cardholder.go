//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// IssuingCardholderRequirementsDisabledReason is the possible values for the disabled reason on an
// issuing cardholder.
type IssuingCardholderRequirementsDisabledReason string

// List of values that IssuingCardholderRequirementsDisabledReason can take.
const (
	IssuingCardholderRequirementsDisabledReasonListed         IssuingCardholderRequirementsDisabledReason = "listed"
	IssuingCardholderRequirementsDisabledReasonRejectedListed IssuingCardholderRequirementsDisabledReason = "rejected.listed"
	IssuingCardholderRequirementsDisabledReasonUnderReview    IssuingCardholderRequirementsDisabledReason = "under_review"
)

type IssuingCardholderRequirementsPastDue string

const (
	IssuingCardholderRequirementsPastDueCompanyTaxID                   IssuingCardholderRequirementsPastDue = "company.tax_id"
	IssuingCardholderRequirementsPastDueIndividualDobDay               IssuingCardholderRequirementsPastDue = "individual.dob.day"
	IssuingCardholderRequirementsPastDueIndividualDobMonth             IssuingCardholderRequirementsPastDue = "individual.dob.month"
	IssuingCardholderRequirementsPastDueIndividualDobYear              IssuingCardholderRequirementsPastDue = "individual.dob.year"
	IssuingCardholderRequirementsPastDueIndividualFirstName            IssuingCardholderRequirementsPastDue = "individual.first_name"
	IssuingCardholderRequirementsPastDueIndividualLastName             IssuingCardholderRequirementsPastDue = "individual.last_name"
	IssuingCardholderRequirementsPastDueIndividualVerificationDocument IssuingCardholderRequirementsPastDue = "individual.verification.document"
)

type IssuingCardholderSpendingControlsAllowedCategory string

const (
	IssuingCardholderSpendingControlsAllowedCategoryAcRefrigerationRepair                                      IssuingCardholderSpendingControlsAllowedCategory = "ac_refrigeration_repair"
	IssuingCardholderSpendingControlsAllowedCategoryAccountingBookkeepingServices                              IssuingCardholderSpendingControlsAllowedCategory = "accounting_bookkeeping_services"
	IssuingCardholderSpendingControlsAllowedCategoryAdvertisingServices                                        IssuingCardholderSpendingControlsAllowedCategory = "advertising_services"
	IssuingCardholderSpendingControlsAllowedCategoryAgriculturalCooperative                                    IssuingCardholderSpendingControlsAllowedCategory = "agricultural_cooperative"
	IssuingCardholderSpendingControlsAllowedCategoryAirlinesAirCarriers                                        IssuingCardholderSpendingControlsAllowedCategory = "airlines_air_carriers"
	IssuingCardholderSpendingControlsAllowedCategoryAirportsFlyingFields                                       IssuingCardholderSpendingControlsAllowedCategory = "airports_flying_fields"
	IssuingCardholderSpendingControlsAllowedCategoryAmbulanceServices                                          IssuingCardholderSpendingControlsAllowedCategory = "ambulance_services"
	IssuingCardholderSpendingControlsAllowedCategoryAmusementParksCarnivals                                    IssuingCardholderSpendingControlsAllowedCategory = "amusement_parks_carnivals"
	IssuingCardholderSpendingControlsAllowedCategoryAntiqueReproductions                                       IssuingCardholderSpendingControlsAllowedCategory = "antique_reproductions"
	IssuingCardholderSpendingControlsAllowedCategoryAntiqueShops                                               IssuingCardholderSpendingControlsAllowedCategory = "antique_shops"
	IssuingCardholderSpendingControlsAllowedCategoryAquariums                                                  IssuingCardholderSpendingControlsAllowedCategory = "aquariums"
	IssuingCardholderSpendingControlsAllowedCategoryArchitecturalSurveyingServices                             IssuingCardholderSpendingControlsAllowedCategory = "architectural_surveying_services"
	IssuingCardholderSpendingControlsAllowedCategoryArtDealersAndGalleries                                     IssuingCardholderSpendingControlsAllowedCategory = "art_dealers_and_galleries"
	IssuingCardholderSpendingControlsAllowedCategoryArtistsSupplyAndCraftShops                                 IssuingCardholderSpendingControlsAllowedCategory = "artists_supply_and_craft_shops"
	IssuingCardholderSpendingControlsAllowedCategoryAutoAndHomeSupplyStores                                    IssuingCardholderSpendingControlsAllowedCategory = "auto_and_home_supply_stores"
	IssuingCardholderSpendingControlsAllowedCategoryAutoBodyRepairShops                                        IssuingCardholderSpendingControlsAllowedCategory = "auto_body_repair_shops"
	IssuingCardholderSpendingControlsAllowedCategoryAutoPaintShops                                             IssuingCardholderSpendingControlsAllowedCategory = "auto_paint_shops"
	IssuingCardholderSpendingControlsAllowedCategoryAutoServiceShops                                           IssuingCardholderSpendingControlsAllowedCategory = "auto_service_shops"
	IssuingCardholderSpendingControlsAllowedCategoryAutomatedCashDisburse                                      IssuingCardholderSpendingControlsAllowedCategory = "automated_cash_disburse"
	IssuingCardholderSpendingControlsAllowedCategoryAutomatedFuelDispensers                                    IssuingCardholderSpendingControlsAllowedCategory = "automated_fuel_dispensers"
	IssuingCardholderSpendingControlsAllowedCategoryAutomobileAssociations                                     IssuingCardholderSpendingControlsAllowedCategory = "automobile_associations"
	IssuingCardholderSpendingControlsAllowedCategoryAutomotivePartsAndAccessoriesStores                        IssuingCardholderSpendingControlsAllowedCategory = "automotive_parts_and_accessories_stores"
	IssuingCardholderSpendingControlsAllowedCategoryAutomotiveTireStores                                       IssuingCardholderSpendingControlsAllowedCategory = "automotive_tire_stores"
	IssuingCardholderSpendingControlsAllowedCategoryBailAndBondPayments                                        IssuingCardholderSpendingControlsAllowedCategory = "bail_and_bond_payments"
	IssuingCardholderSpendingControlsAllowedCategoryBakeries                                                   IssuingCardholderSpendingControlsAllowedCategory = "bakeries"
	IssuingCardholderSpendingControlsAllowedCategoryBandsOrchestras                                            IssuingCardholderSpendingControlsAllowedCategory = "bands_orchestras"
	IssuingCardholderSpendingControlsAllowedCategoryBarberAndBeautyShops                                       IssuingCardholderSpendingControlsAllowedCategory = "barber_and_beauty_shops"
	IssuingCardholderSpendingControlsAllowedCategoryBettingCasinoGambling                                      IssuingCardholderSpendingControlsAllowedCategory = "betting_casino_gambling"
	IssuingCardholderSpendingControlsAllowedCategoryBicycleShops                                               IssuingCardholderSpendingControlsAllowedCategory = "bicycle_shops"
	IssuingCardholderSpendingControlsAllowedCategoryBilliardPoolEstablishments                                 IssuingCardholderSpendingControlsAllowedCategory = "billiard_pool_establishments"
	IssuingCardholderSpendingControlsAllowedCategoryBoatDealers                                                IssuingCardholderSpendingControlsAllowedCategory = "boat_dealers"
	IssuingCardholderSpendingControlsAllowedCategoryBoatRentalsAndLeases                                       IssuingCardholderSpendingControlsAllowedCategory = "boat_rentals_and_leases"
	IssuingCardholderSpendingControlsAllowedCategoryBookStores                                                 IssuingCardholderSpendingControlsAllowedCategory = "book_stores"
	IssuingCardholderSpendingControlsAllowedCategoryBooksPeriodicalsAndNewspapers                              IssuingCardholderSpendingControlsAllowedCategory = "books_periodicals_and_newspapers"
	IssuingCardholderSpendingControlsAllowedCategoryBowlingAlleys                                              IssuingCardholderSpendingControlsAllowedCategory = "bowling_alleys"
	IssuingCardholderSpendingControlsAllowedCategoryBusLines                                                   IssuingCardholderSpendingControlsAllowedCategory = "bus_lines"
	IssuingCardholderSpendingControlsAllowedCategoryBusinessSecretarialSchools                                 IssuingCardholderSpendingControlsAllowedCategory = "business_secretarial_schools"
	IssuingCardholderSpendingControlsAllowedCategoryBuyingShoppingServices                                     IssuingCardholderSpendingControlsAllowedCategory = "buying_shopping_services"
	IssuingCardholderSpendingControlsAllowedCategoryCableSatelliteAndOtherPayTelevisionAndRadio                IssuingCardholderSpendingControlsAllowedCategory = "cable_satellite_and_other_pay_television_and_radio"
	IssuingCardholderSpendingControlsAllowedCategoryCameraAndPhotographicSupplyStores                          IssuingCardholderSpendingControlsAllowedCategory = "camera_and_photographic_supply_stores"
	IssuingCardholderSpendingControlsAllowedCategoryCandyNutAndConfectioneryStores                             IssuingCardholderSpendingControlsAllowedCategory = "candy_nut_and_confectionery_stores"
	IssuingCardholderSpendingControlsAllowedCategoryCarAndTruckDealersNewUsed                                  IssuingCardholderSpendingControlsAllowedCategory = "car_and_truck_dealers_new_used"
	IssuingCardholderSpendingControlsAllowedCategoryCarAndTruckDealersUsedOnly                                 IssuingCardholderSpendingControlsAllowedCategory = "car_and_truck_dealers_used_only"
	IssuingCardholderSpendingControlsAllowedCategoryCarRentalAgencies                                          IssuingCardholderSpendingControlsAllowedCategory = "car_rental_agencies"
	IssuingCardholderSpendingControlsAllowedCategoryCarWashes                                                  IssuingCardholderSpendingControlsAllowedCategory = "car_washes"
	IssuingCardholderSpendingControlsAllowedCategoryCarpentryServices                                          IssuingCardholderSpendingControlsAllowedCategory = "carpentry_services"
	IssuingCardholderSpendingControlsAllowedCategoryCarpetUpholsteryCleaning                                   IssuingCardholderSpendingControlsAllowedCategory = "carpet_upholstery_cleaning"
	IssuingCardholderSpendingControlsAllowedCategoryCaterers                                                   IssuingCardholderSpendingControlsAllowedCategory = "caterers"
	IssuingCardholderSpendingControlsAllowedCategoryCharitableAndSocialServiceOrganizationsFundraising         IssuingCardholderSpendingControlsAllowedCategory = "charitable_and_social_service_organizations_fundraising"
	IssuingCardholderSpendingControlsAllowedCategoryChemicalsAndAlliedProducts                                 IssuingCardholderSpendingControlsAllowedCategory = "chemicals_and_allied_products"
	IssuingCardholderSpendingControlsAllowedCategoryChildCareServices                                          IssuingCardholderSpendingControlsAllowedCategory = "child_care_services"
	IssuingCardholderSpendingControlsAllowedCategoryChildrensAndInfantsWearStores                              IssuingCardholderSpendingControlsAllowedCategory = "childrens_and_infants_wear_stores"
	IssuingCardholderSpendingControlsAllowedCategoryChiropodistsPodiatrists                                    IssuingCardholderSpendingControlsAllowedCategory = "chiropodists_podiatrists"
	IssuingCardholderSpendingControlsAllowedCategoryChiropractors                                              IssuingCardholderSpendingControlsAllowedCategory = "chiropractors"
	IssuingCardholderSpendingControlsAllowedCategoryCigarStoresAndStands                                       IssuingCardholderSpendingControlsAllowedCategory = "cigar_stores_and_stands"
	IssuingCardholderSpendingControlsAllowedCategoryCivicSocialFraternalAssociations                           IssuingCardholderSpendingControlsAllowedCategory = "civic_social_fraternal_associations"
	IssuingCardholderSpendingControlsAllowedCategoryCleaningAndMaintenance                                     IssuingCardholderSpendingControlsAllowedCategory = "cleaning_and_maintenance"
	IssuingCardholderSpendingControlsAllowedCategoryClothingRental                                             IssuingCardholderSpendingControlsAllowedCategory = "clothing_rental"
	IssuingCardholderSpendingControlsAllowedCategoryCollegesUniversities                                       IssuingCardholderSpendingControlsAllowedCategory = "colleges_universities"
	IssuingCardholderSpendingControlsAllowedCategoryCommercialEquipment                                        IssuingCardholderSpendingControlsAllowedCategory = "commercial_equipment"
	IssuingCardholderSpendingControlsAllowedCategoryCommercialFootwear                                         IssuingCardholderSpendingControlsAllowedCategory = "commercial_footwear"
	IssuingCardholderSpendingControlsAllowedCategoryCommercialPhotographyArtAndGraphics                        IssuingCardholderSpendingControlsAllowedCategory = "commercial_photography_art_and_graphics"
	IssuingCardholderSpendingControlsAllowedCategoryCommuterTransportAndFerries                                IssuingCardholderSpendingControlsAllowedCategory = "commuter_transport_and_ferries"
	IssuingCardholderSpendingControlsAllowedCategoryComputerNetworkServices                                    IssuingCardholderSpendingControlsAllowedCategory = "computer_network_services"
	IssuingCardholderSpendingControlsAllowedCategoryComputerProgramming                                        IssuingCardholderSpendingControlsAllowedCategory = "computer_programming"
	IssuingCardholderSpendingControlsAllowedCategoryComputerRepair                                             IssuingCardholderSpendingControlsAllowedCategory = "computer_repair"
	IssuingCardholderSpendingControlsAllowedCategoryComputerSoftwareStores                                     IssuingCardholderSpendingControlsAllowedCategory = "computer_software_stores"
	IssuingCardholderSpendingControlsAllowedCategoryComputersPeripheralsAndSoftware                            IssuingCardholderSpendingControlsAllowedCategory = "computers_peripherals_and_software"
	IssuingCardholderSpendingControlsAllowedCategoryConcreteWorkServices                                       IssuingCardholderSpendingControlsAllowedCategory = "concrete_work_services"
	IssuingCardholderSpendingControlsAllowedCategoryConstructionMaterials                                      IssuingCardholderSpendingControlsAllowedCategory = "construction_materials"
	IssuingCardholderSpendingControlsAllowedCategoryConsultingPublicRelations                                  IssuingCardholderSpendingControlsAllowedCategory = "consulting_public_relations"
	IssuingCardholderSpendingControlsAllowedCategoryCorrespondenceSchools                                      IssuingCardholderSpendingControlsAllowedCategory = "correspondence_schools"
	IssuingCardholderSpendingControlsAllowedCategoryCosmeticStores                                             IssuingCardholderSpendingControlsAllowedCategory = "cosmetic_stores"
	IssuingCardholderSpendingControlsAllowedCategoryCounselingServices                                         IssuingCardholderSpendingControlsAllowedCategory = "counseling_services"
	IssuingCardholderSpendingControlsAllowedCategoryCountryClubs                                               IssuingCardholderSpendingControlsAllowedCategory = "country_clubs"
	IssuingCardholderSpendingControlsAllowedCategoryCourierServices                                            IssuingCardholderSpendingControlsAllowedCategory = "courier_services"
	IssuingCardholderSpendingControlsAllowedCategoryCourtCosts                                                 IssuingCardholderSpendingControlsAllowedCategory = "court_costs"
	IssuingCardholderSpendingControlsAllowedCategoryCreditReportingAgencies                                    IssuingCardholderSpendingControlsAllowedCategory = "credit_reporting_agencies"
	IssuingCardholderSpendingControlsAllowedCategoryCruiseLines                                                IssuingCardholderSpendingControlsAllowedCategory = "cruise_lines"
	IssuingCardholderSpendingControlsAllowedCategoryDairyProductsStores                                        IssuingCardholderSpendingControlsAllowedCategory = "dairy_products_stores"
	IssuingCardholderSpendingControlsAllowedCategoryDanceHallStudiosSchools                                    IssuingCardholderSpendingControlsAllowedCategory = "dance_hall_studios_schools"
	IssuingCardholderSpendingControlsAllowedCategoryDatingEscortServices                                       IssuingCardholderSpendingControlsAllowedCategory = "dating_escort_services"
	IssuingCardholderSpendingControlsAllowedCategoryDentistsOrthodontists                                      IssuingCardholderSpendingControlsAllowedCategory = "dentists_orthodontists"
	IssuingCardholderSpendingControlsAllowedCategoryDepartmentStores                                           IssuingCardholderSpendingControlsAllowedCategory = "department_stores"
	IssuingCardholderSpendingControlsAllowedCategoryDetectiveAgencies                                          IssuingCardholderSpendingControlsAllowedCategory = "detective_agencies"
	IssuingCardholderSpendingControlsAllowedCategoryDigitalGoodsApplications                                   IssuingCardholderSpendingControlsAllowedCategory = "digital_goods_applications"
	IssuingCardholderSpendingControlsAllowedCategoryDigitalGoodsGames                                          IssuingCardholderSpendingControlsAllowedCategory = "digital_goods_games"
	IssuingCardholderSpendingControlsAllowedCategoryDigitalGoodsLargeVolume                                    IssuingCardholderSpendingControlsAllowedCategory = "digital_goods_large_volume"
	IssuingCardholderSpendingControlsAllowedCategoryDigitalGoodsMedia                                          IssuingCardholderSpendingControlsAllowedCategory = "digital_goods_media"
	IssuingCardholderSpendingControlsAllowedCategoryDirectMarketingCatalogMerchant                             IssuingCardholderSpendingControlsAllowedCategory = "direct_marketing_catalog_merchant"
	IssuingCardholderSpendingControlsAllowedCategoryDirectMarketingCombinationCatalogAndRetailMerchant         IssuingCardholderSpendingControlsAllowedCategory = "direct_marketing_combination_catalog_and_retail_merchant"
	IssuingCardholderSpendingControlsAllowedCategoryDirectMarketingInboundTelemarketing                        IssuingCardholderSpendingControlsAllowedCategory = "direct_marketing_inbound_telemarketing"
	IssuingCardholderSpendingControlsAllowedCategoryDirectMarketingInsuranceServices                           IssuingCardholderSpendingControlsAllowedCategory = "direct_marketing_insurance_services"
	IssuingCardholderSpendingControlsAllowedCategoryDirectMarketingOther                                       IssuingCardholderSpendingControlsAllowedCategory = "direct_marketing_other"
	IssuingCardholderSpendingControlsAllowedCategoryDirectMarketingOutboundTelemarketing                       IssuingCardholderSpendingControlsAllowedCategory = "direct_marketing_outbound_telemarketing"
	IssuingCardholderSpendingControlsAllowedCategoryDirectMarketingSubscription                                IssuingCardholderSpendingControlsAllowedCategory = "direct_marketing_subscription"
	IssuingCardholderSpendingControlsAllowedCategoryDirectMarketingTravel                                      IssuingCardholderSpendingControlsAllowedCategory = "direct_marketing_travel"
	IssuingCardholderSpendingControlsAllowedCategoryDiscountStores                                             IssuingCardholderSpendingControlsAllowedCategory = "discount_stores"
	IssuingCardholderSpendingControlsAllowedCategoryDoctors                                                    IssuingCardholderSpendingControlsAllowedCategory = "doctors"
	IssuingCardholderSpendingControlsAllowedCategoryDoorToDoorSales                                            IssuingCardholderSpendingControlsAllowedCategory = "door_to_door_sales"
	IssuingCardholderSpendingControlsAllowedCategoryDraperyWindowCoveringAndUpholsteryStores                   IssuingCardholderSpendingControlsAllowedCategory = "drapery_window_covering_and_upholstery_stores"
	IssuingCardholderSpendingControlsAllowedCategoryDrinkingPlaces                                             IssuingCardholderSpendingControlsAllowedCategory = "drinking_places"
	IssuingCardholderSpendingControlsAllowedCategoryDrugStoresAndPharmacies                                    IssuingCardholderSpendingControlsAllowedCategory = "drug_stores_and_pharmacies"
	IssuingCardholderSpendingControlsAllowedCategoryDrugsDrugProprietariesAndDruggistSundries                  IssuingCardholderSpendingControlsAllowedCategory = "drugs_drug_proprietaries_and_druggist_sundries"
	IssuingCardholderSpendingControlsAllowedCategoryDryCleaners                                                IssuingCardholderSpendingControlsAllowedCategory = "dry_cleaners"
	IssuingCardholderSpendingControlsAllowedCategoryDurableGoods                                               IssuingCardholderSpendingControlsAllowedCategory = "durable_goods"
	IssuingCardholderSpendingControlsAllowedCategoryDutyFreeStores                                             IssuingCardholderSpendingControlsAllowedCategory = "duty_free_stores"
	IssuingCardholderSpendingControlsAllowedCategoryEatingPlacesRestaurants                                    IssuingCardholderSpendingControlsAllowedCategory = "eating_places_restaurants"
	IssuingCardholderSpendingControlsAllowedCategoryEducationalServices                                        IssuingCardholderSpendingControlsAllowedCategory = "educational_services"
	IssuingCardholderSpendingControlsAllowedCategoryElectricRazorStores                                        IssuingCardholderSpendingControlsAllowedCategory = "electric_razor_stores"
	IssuingCardholderSpendingControlsAllowedCategoryElectricalPartsAndEquipment                                IssuingCardholderSpendingControlsAllowedCategory = "electrical_parts_and_equipment"
	IssuingCardholderSpendingControlsAllowedCategoryElectricalServices                                         IssuingCardholderSpendingControlsAllowedCategory = "electrical_services"
	IssuingCardholderSpendingControlsAllowedCategoryElectronicsRepairShops                                     IssuingCardholderSpendingControlsAllowedCategory = "electronics_repair_shops"
	IssuingCardholderSpendingControlsAllowedCategoryElectronicsStores                                          IssuingCardholderSpendingControlsAllowedCategory = "electronics_stores"
	IssuingCardholderSpendingControlsAllowedCategoryElementarySecondarySchools                                 IssuingCardholderSpendingControlsAllowedCategory = "elementary_secondary_schools"
	IssuingCardholderSpendingControlsAllowedCategoryEmploymentTempAgencies                                     IssuingCardholderSpendingControlsAllowedCategory = "employment_temp_agencies"
	IssuingCardholderSpendingControlsAllowedCategoryEquipmentRental                                            IssuingCardholderSpendingControlsAllowedCategory = "equipment_rental"
	IssuingCardholderSpendingControlsAllowedCategoryExterminatingServices                                      IssuingCardholderSpendingControlsAllowedCategory = "exterminating_services"
	IssuingCardholderSpendingControlsAllowedCategoryFamilyClothingStores                                       IssuingCardholderSpendingControlsAllowedCategory = "family_clothing_stores"
	IssuingCardholderSpendingControlsAllowedCategoryFastFoodRestaurants                                        IssuingCardholderSpendingControlsAllowedCategory = "fast_food_restaurants"
	IssuingCardholderSpendingControlsAllowedCategoryFinancialInstitutions                                      IssuingCardholderSpendingControlsAllowedCategory = "financial_institutions"
	IssuingCardholderSpendingControlsAllowedCategoryFinesGovernmentAdministrativeEntities                      IssuingCardholderSpendingControlsAllowedCategory = "fines_government_administrative_entities"
	IssuingCardholderSpendingControlsAllowedCategoryFireplaceFireplaceScreensAndAccessoriesStores              IssuingCardholderSpendingControlsAllowedCategory = "fireplace_fireplace_screens_and_accessories_stores"
	IssuingCardholderSpendingControlsAllowedCategoryFloorCoveringStores                                        IssuingCardholderSpendingControlsAllowedCategory = "floor_covering_stores"
	IssuingCardholderSpendingControlsAllowedCategoryFlorists                                                   IssuingCardholderSpendingControlsAllowedCategory = "florists"
	IssuingCardholderSpendingControlsAllowedCategoryFloristsSuppliesNurseryStockAndFlowers                     IssuingCardholderSpendingControlsAllowedCategory = "florists_supplies_nursery_stock_and_flowers"
	IssuingCardholderSpendingControlsAllowedCategoryFreezerAndLockerMeatProvisioners                           IssuingCardholderSpendingControlsAllowedCategory = "freezer_and_locker_meat_provisioners"
	IssuingCardholderSpendingControlsAllowedCategoryFuelDealersNonAutomotive                                   IssuingCardholderSpendingControlsAllowedCategory = "fuel_dealers_non_automotive"
	IssuingCardholderSpendingControlsAllowedCategoryFuneralServicesCrematories                                 IssuingCardholderSpendingControlsAllowedCategory = "funeral_services_crematories"
	IssuingCardholderSpendingControlsAllowedCategoryFurnitureHomeFurnishingsAndEquipmentStoresExceptAppliances IssuingCardholderSpendingControlsAllowedCategory = "furniture_home_furnishings_and_equipment_stores_except_appliances"
	IssuingCardholderSpendingControlsAllowedCategoryFurnitureRepairRefinishing                                 IssuingCardholderSpendingControlsAllowedCategory = "furniture_repair_refinishing"
	IssuingCardholderSpendingControlsAllowedCategoryFurriersAndFurShops                                        IssuingCardholderSpendingControlsAllowedCategory = "furriers_and_fur_shops"
	IssuingCardholderSpendingControlsAllowedCategoryGeneralServices                                            IssuingCardholderSpendingControlsAllowedCategory = "general_services"
	IssuingCardholderSpendingControlsAllowedCategoryGiftCardNoveltyAndSouvenirShops                            IssuingCardholderSpendingControlsAllowedCategory = "gift_card_novelty_and_souvenir_shops"
	IssuingCardholderSpendingControlsAllowedCategoryGlassPaintAndWallpaperStores                               IssuingCardholderSpendingControlsAllowedCategory = "glass_paint_and_wallpaper_stores"
	IssuingCardholderSpendingControlsAllowedCategoryGlasswareCrystalStores                                     IssuingCardholderSpendingControlsAllowedCategory = "glassware_crystal_stores"
	IssuingCardholderSpendingControlsAllowedCategoryGolfCoursesPublic                                          IssuingCardholderSpendingControlsAllowedCategory = "golf_courses_public"
	IssuingCardholderSpendingControlsAllowedCategoryGovernmentServices                                         IssuingCardholderSpendingControlsAllowedCategory = "government_services"
	IssuingCardholderSpendingControlsAllowedCategoryGroceryStoresSupermarkets                                  IssuingCardholderSpendingControlsAllowedCategory = "grocery_stores_supermarkets"
	IssuingCardholderSpendingControlsAllowedCategoryHardwareEquipmentAndSupplies                               IssuingCardholderSpendingControlsAllowedCategory = "hardware_equipment_and_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryHardwareStores                                             IssuingCardholderSpendingControlsAllowedCategory = "hardware_stores"
	IssuingCardholderSpendingControlsAllowedCategoryHealthAndBeautySpas                                        IssuingCardholderSpendingControlsAllowedCategory = "health_and_beauty_spas"
	IssuingCardholderSpendingControlsAllowedCategoryHearingAidsSalesAndSupplies                                IssuingCardholderSpendingControlsAllowedCategory = "hearing_aids_sales_and_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryHeatingPlumbingAC                                          IssuingCardholderSpendingControlsAllowedCategory = "heating_plumbing_a_c"
	IssuingCardholderSpendingControlsAllowedCategoryHobbyToyAndGameShops                                       IssuingCardholderSpendingControlsAllowedCategory = "hobby_toy_and_game_shops"
	IssuingCardholderSpendingControlsAllowedCategoryHomeSupplyWarehouseStores                                  IssuingCardholderSpendingControlsAllowedCategory = "home_supply_warehouse_stores"
	IssuingCardholderSpendingControlsAllowedCategoryHospitals                                                  IssuingCardholderSpendingControlsAllowedCategory = "hospitals"
	IssuingCardholderSpendingControlsAllowedCategoryHotelsMotelsAndResorts                                     IssuingCardholderSpendingControlsAllowedCategory = "hotels_motels_and_resorts"
	IssuingCardholderSpendingControlsAllowedCategoryHouseholdApplianceStores                                   IssuingCardholderSpendingControlsAllowedCategory = "household_appliance_stores"
	IssuingCardholderSpendingControlsAllowedCategoryIndustrialSupplies                                         IssuingCardholderSpendingControlsAllowedCategory = "industrial_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryInformationRetrievalServices                               IssuingCardholderSpendingControlsAllowedCategory = "information_retrieval_services"
	IssuingCardholderSpendingControlsAllowedCategoryInsuranceDefault                                           IssuingCardholderSpendingControlsAllowedCategory = "insurance_default"
	IssuingCardholderSpendingControlsAllowedCategoryInsuranceUnderwritingPremiums                              IssuingCardholderSpendingControlsAllowedCategory = "insurance_underwriting_premiums"
	IssuingCardholderSpendingControlsAllowedCategoryIntraCompanyPurchases                                      IssuingCardholderSpendingControlsAllowedCategory = "intra_company_purchases"
	IssuingCardholderSpendingControlsAllowedCategoryJewelryStoresWatchesClocksAndSilverwareStores              IssuingCardholderSpendingControlsAllowedCategory = "jewelry_stores_watches_clocks_and_silverware_stores"
	IssuingCardholderSpendingControlsAllowedCategoryLandscapingServices                                        IssuingCardholderSpendingControlsAllowedCategory = "landscaping_services"
	IssuingCardholderSpendingControlsAllowedCategoryLaundries                                                  IssuingCardholderSpendingControlsAllowedCategory = "laundries"
	IssuingCardholderSpendingControlsAllowedCategoryLaundryCleaningServices                                    IssuingCardholderSpendingControlsAllowedCategory = "laundry_cleaning_services"
	IssuingCardholderSpendingControlsAllowedCategoryLegalServicesAttorneys                                     IssuingCardholderSpendingControlsAllowedCategory = "legal_services_attorneys"
	IssuingCardholderSpendingControlsAllowedCategoryLuggageAndLeatherGoodsStores                               IssuingCardholderSpendingControlsAllowedCategory = "luggage_and_leather_goods_stores"
	IssuingCardholderSpendingControlsAllowedCategoryLumberBuildingMaterialsStores                              IssuingCardholderSpendingControlsAllowedCategory = "lumber_building_materials_stores"
	IssuingCardholderSpendingControlsAllowedCategoryManualCashDisburse                                         IssuingCardholderSpendingControlsAllowedCategory = "manual_cash_disburse"
	IssuingCardholderSpendingControlsAllowedCategoryMarinasServiceAndSupplies                                  IssuingCardholderSpendingControlsAllowedCategory = "marinas_service_and_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryMasonryStoneworkAndPlaster                                 IssuingCardholderSpendingControlsAllowedCategory = "masonry_stonework_and_plaster"
	IssuingCardholderSpendingControlsAllowedCategoryMassageParlors                                             IssuingCardholderSpendingControlsAllowedCategory = "massage_parlors"
	IssuingCardholderSpendingControlsAllowedCategoryMedicalAndDentalLabs                                       IssuingCardholderSpendingControlsAllowedCategory = "medical_and_dental_labs"
	IssuingCardholderSpendingControlsAllowedCategoryMedicalDentalOphthalmicAndHospitalEquipmentAndSupplies     IssuingCardholderSpendingControlsAllowedCategory = "medical_dental_ophthalmic_and_hospital_equipment_and_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryMedicalServices                                            IssuingCardholderSpendingControlsAllowedCategory = "medical_services"
	IssuingCardholderSpendingControlsAllowedCategoryMembershipOrganizations                                    IssuingCardholderSpendingControlsAllowedCategory = "membership_organizations"
	IssuingCardholderSpendingControlsAllowedCategoryMensAndBoysClothingAndAccessoriesStores                    IssuingCardholderSpendingControlsAllowedCategory = "mens_and_boys_clothing_and_accessories_stores"
	IssuingCardholderSpendingControlsAllowedCategoryMensWomensClothingStores                                   IssuingCardholderSpendingControlsAllowedCategory = "mens_womens_clothing_stores"
	IssuingCardholderSpendingControlsAllowedCategoryMetalServiceCenters                                        IssuingCardholderSpendingControlsAllowedCategory = "metal_service_centers"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneous                                              IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousApparelAndAccessoryShops                      IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_apparel_and_accessory_shops"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousAutoDealers                                   IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_auto_dealers"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousBusinessServices                              IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_business_services"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousFoodStores                                    IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_food_stores"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousGeneralMerchandise                            IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_general_merchandise"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousGeneralServices                               IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_general_services"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousHomeFurnishingSpecialtyStores                 IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_home_furnishing_specialty_stores"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousPublishingAndPrinting                         IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_publishing_and_printing"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousRecreationServices                            IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_recreation_services"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousRepairShops                                   IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_repair_shops"
	IssuingCardholderSpendingControlsAllowedCategoryMiscellaneousSpecialtyRetail                               IssuingCardholderSpendingControlsAllowedCategory = "miscellaneous_specialty_retail"
	IssuingCardholderSpendingControlsAllowedCategoryMobileHomeDealers                                          IssuingCardholderSpendingControlsAllowedCategory = "mobile_home_dealers"
	IssuingCardholderSpendingControlsAllowedCategoryMotionPictureTheaters                                      IssuingCardholderSpendingControlsAllowedCategory = "motion_picture_theaters"
	IssuingCardholderSpendingControlsAllowedCategoryMotorFreightCarriersAndTrucking                            IssuingCardholderSpendingControlsAllowedCategory = "motor_freight_carriers_and_trucking"
	IssuingCardholderSpendingControlsAllowedCategoryMotorHomesDealers                                          IssuingCardholderSpendingControlsAllowedCategory = "motor_homes_dealers"
	IssuingCardholderSpendingControlsAllowedCategoryMotorVehicleSuppliesAndNewParts                            IssuingCardholderSpendingControlsAllowedCategory = "motor_vehicle_supplies_and_new_parts"
	IssuingCardholderSpendingControlsAllowedCategoryMotorcycleShopsAndDealers                                  IssuingCardholderSpendingControlsAllowedCategory = "motorcycle_shops_and_dealers"
	IssuingCardholderSpendingControlsAllowedCategoryMotorcycleShopsDealers                                     IssuingCardholderSpendingControlsAllowedCategory = "motorcycle_shops_dealers"
	IssuingCardholderSpendingControlsAllowedCategoryMusicStoresMusicalInstrumentsPianosAndSheetMusic           IssuingCardholderSpendingControlsAllowedCategory = "music_stores_musical_instruments_pianos_and_sheet_music"
	IssuingCardholderSpendingControlsAllowedCategoryNewsDealersAndNewsstands                                   IssuingCardholderSpendingControlsAllowedCategory = "news_dealers_and_newsstands"
	IssuingCardholderSpendingControlsAllowedCategoryNonFiMoneyOrders                                           IssuingCardholderSpendingControlsAllowedCategory = "non_fi_money_orders"
	IssuingCardholderSpendingControlsAllowedCategoryNonFiStoredValueCardPurchaseLoad                           IssuingCardholderSpendingControlsAllowedCategory = "non_fi_stored_value_card_purchase_load"
	IssuingCardholderSpendingControlsAllowedCategoryNondurableGoods                                            IssuingCardholderSpendingControlsAllowedCategory = "nondurable_goods"
	IssuingCardholderSpendingControlsAllowedCategoryNurseriesLawnAndGardenSupplyStores                         IssuingCardholderSpendingControlsAllowedCategory = "nurseries_lawn_and_garden_supply_stores"
	IssuingCardholderSpendingControlsAllowedCategoryNursingPersonalCare                                        IssuingCardholderSpendingControlsAllowedCategory = "nursing_personal_care"
	IssuingCardholderSpendingControlsAllowedCategoryOfficeAndCommercialFurniture                               IssuingCardholderSpendingControlsAllowedCategory = "office_and_commercial_furniture"
	IssuingCardholderSpendingControlsAllowedCategoryOpticiansEyeglasses                                        IssuingCardholderSpendingControlsAllowedCategory = "opticians_eyeglasses"
	IssuingCardholderSpendingControlsAllowedCategoryOptometristsOphthalmologist                                IssuingCardholderSpendingControlsAllowedCategory = "optometrists_ophthalmologist"
	IssuingCardholderSpendingControlsAllowedCategoryOrthopedicGoodsProstheticDevices                           IssuingCardholderSpendingControlsAllowedCategory = "orthopedic_goods_prosthetic_devices"
	IssuingCardholderSpendingControlsAllowedCategoryOsteopaths                                                 IssuingCardholderSpendingControlsAllowedCategory = "osteopaths"
	IssuingCardholderSpendingControlsAllowedCategoryPackageStoresBeerWineAndLiquor                             IssuingCardholderSpendingControlsAllowedCategory = "package_stores_beer_wine_and_liquor"
	IssuingCardholderSpendingControlsAllowedCategoryPaintsVarnishesAndSupplies                                 IssuingCardholderSpendingControlsAllowedCategory = "paints_varnishes_and_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryParkingLotsGarages                                         IssuingCardholderSpendingControlsAllowedCategory = "parking_lots_garages"
	IssuingCardholderSpendingControlsAllowedCategoryPassengerRailways                                          IssuingCardholderSpendingControlsAllowedCategory = "passenger_railways"
	IssuingCardholderSpendingControlsAllowedCategoryPawnShops                                                  IssuingCardholderSpendingControlsAllowedCategory = "pawn_shops"
	IssuingCardholderSpendingControlsAllowedCategoryPetShopsPetFoodAndSupplies                                 IssuingCardholderSpendingControlsAllowedCategory = "pet_shops_pet_food_and_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryPetroleumAndPetroleumProducts                              IssuingCardholderSpendingControlsAllowedCategory = "petroleum_and_petroleum_products"
	IssuingCardholderSpendingControlsAllowedCategoryPhotoDeveloping                                            IssuingCardholderSpendingControlsAllowedCategory = "photo_developing"
	IssuingCardholderSpendingControlsAllowedCategoryPhotographicPhotocopyMicrofilmEquipmentAndSupplies         IssuingCardholderSpendingControlsAllowedCategory = "photographic_photocopy_microfilm_equipment_and_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryPhotographicStudios                                        IssuingCardholderSpendingControlsAllowedCategory = "photographic_studios"
	IssuingCardholderSpendingControlsAllowedCategoryPictureVideoProduction                                     IssuingCardholderSpendingControlsAllowedCategory = "picture_video_production"
	IssuingCardholderSpendingControlsAllowedCategoryPieceGoodsNotionsAndOtherDryGoods                          IssuingCardholderSpendingControlsAllowedCategory = "piece_goods_notions_and_other_dry_goods"
	IssuingCardholderSpendingControlsAllowedCategoryPlumbingHeatingEquipmentAndSupplies                        IssuingCardholderSpendingControlsAllowedCategory = "plumbing_heating_equipment_and_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryPoliticalOrganizations                                     IssuingCardholderSpendingControlsAllowedCategory = "political_organizations"
	IssuingCardholderSpendingControlsAllowedCategoryPostalServicesGovernmentOnly                               IssuingCardholderSpendingControlsAllowedCategory = "postal_services_government_only"
	IssuingCardholderSpendingControlsAllowedCategoryPreciousStonesAndMetalsWatchesAndJewelry                   IssuingCardholderSpendingControlsAllowedCategory = "precious_stones_and_metals_watches_and_jewelry"
	IssuingCardholderSpendingControlsAllowedCategoryProfessionalServices                                       IssuingCardholderSpendingControlsAllowedCategory = "professional_services"
	IssuingCardholderSpendingControlsAllowedCategoryPublicWarehousingAndStorage                                IssuingCardholderSpendingControlsAllowedCategory = "public_warehousing_and_storage"
	IssuingCardholderSpendingControlsAllowedCategoryQuickCopyReproAndBlueprint                                 IssuingCardholderSpendingControlsAllowedCategory = "quick_copy_repro_and_blueprint"
	IssuingCardholderSpendingControlsAllowedCategoryRailroads                                                  IssuingCardholderSpendingControlsAllowedCategory = "railroads"
	IssuingCardholderSpendingControlsAllowedCategoryRealEstateAgentsAndManagersRentals                         IssuingCardholderSpendingControlsAllowedCategory = "real_estate_agents_and_managers_rentals"
	IssuingCardholderSpendingControlsAllowedCategoryRecordStores                                               IssuingCardholderSpendingControlsAllowedCategory = "record_stores"
	IssuingCardholderSpendingControlsAllowedCategoryRecreationalVehicleRentals                                 IssuingCardholderSpendingControlsAllowedCategory = "recreational_vehicle_rentals"
	IssuingCardholderSpendingControlsAllowedCategoryReligiousGoodsStores                                       IssuingCardholderSpendingControlsAllowedCategory = "religious_goods_stores"
	IssuingCardholderSpendingControlsAllowedCategoryReligiousOrganizations                                     IssuingCardholderSpendingControlsAllowedCategory = "religious_organizations"
	IssuingCardholderSpendingControlsAllowedCategoryRoofingSidingSheetMetal                                    IssuingCardholderSpendingControlsAllowedCategory = "roofing_siding_sheet_metal"
	IssuingCardholderSpendingControlsAllowedCategorySecretarialSupportServices                                 IssuingCardholderSpendingControlsAllowedCategory = "secretarial_support_services"
	IssuingCardholderSpendingControlsAllowedCategorySecurityBrokersDealers                                     IssuingCardholderSpendingControlsAllowedCategory = "security_brokers_dealers"
	IssuingCardholderSpendingControlsAllowedCategoryServiceStations                                            IssuingCardholderSpendingControlsAllowedCategory = "service_stations"
	IssuingCardholderSpendingControlsAllowedCategorySewingNeedleworkFabricAndPieceGoodsStores                  IssuingCardholderSpendingControlsAllowedCategory = "sewing_needlework_fabric_and_piece_goods_stores"
	IssuingCardholderSpendingControlsAllowedCategoryShoeRepairHatCleaning                                      IssuingCardholderSpendingControlsAllowedCategory = "shoe_repair_hat_cleaning"
	IssuingCardholderSpendingControlsAllowedCategoryShoeStores                                                 IssuingCardholderSpendingControlsAllowedCategory = "shoe_stores"
	IssuingCardholderSpendingControlsAllowedCategorySmallApplianceRepair                                       IssuingCardholderSpendingControlsAllowedCategory = "small_appliance_repair"
	IssuingCardholderSpendingControlsAllowedCategorySnowmobileDealers                                          IssuingCardholderSpendingControlsAllowedCategory = "snowmobile_dealers"
	IssuingCardholderSpendingControlsAllowedCategorySpecialTradeServices                                       IssuingCardholderSpendingControlsAllowedCategory = "special_trade_services"
	IssuingCardholderSpendingControlsAllowedCategorySpecialtyCleaning                                          IssuingCardholderSpendingControlsAllowedCategory = "specialty_cleaning"
	IssuingCardholderSpendingControlsAllowedCategorySportingGoodsStores                                        IssuingCardholderSpendingControlsAllowedCategory = "sporting_goods_stores"
	IssuingCardholderSpendingControlsAllowedCategorySportingRecreationCamps                                    IssuingCardholderSpendingControlsAllowedCategory = "sporting_recreation_camps"
	IssuingCardholderSpendingControlsAllowedCategorySportsAndRidingApparelStores                               IssuingCardholderSpendingControlsAllowedCategory = "sports_and_riding_apparel_stores"
	IssuingCardholderSpendingControlsAllowedCategorySportsClubsFields                                          IssuingCardholderSpendingControlsAllowedCategory = "sports_clubs_fields"
	IssuingCardholderSpendingControlsAllowedCategoryStampAndCoinStores                                         IssuingCardholderSpendingControlsAllowedCategory = "stamp_and_coin_stores"
	IssuingCardholderSpendingControlsAllowedCategoryStationaryOfficeSuppliesPrintingAndWritingPaper            IssuingCardholderSpendingControlsAllowedCategory = "stationary_office_supplies_printing_and_writing_paper"
	IssuingCardholderSpendingControlsAllowedCategoryStationeryStoresOfficeAndSchoolSupplyStores                IssuingCardholderSpendingControlsAllowedCategory = "stationery_stores_office_and_school_supply_stores"
	IssuingCardholderSpendingControlsAllowedCategorySwimmingPoolsSales                                         IssuingCardholderSpendingControlsAllowedCategory = "swimming_pools_sales"
	IssuingCardholderSpendingControlsAllowedCategoryTUiTravelGermany                                           IssuingCardholderSpendingControlsAllowedCategory = "t_ui_travel_germany"
	IssuingCardholderSpendingControlsAllowedCategoryTailorsAlterations                                         IssuingCardholderSpendingControlsAllowedCategory = "tailors_alterations"
	IssuingCardholderSpendingControlsAllowedCategoryTaxPaymentsGovernmentAgencies                              IssuingCardholderSpendingControlsAllowedCategory = "tax_payments_government_agencies"
	IssuingCardholderSpendingControlsAllowedCategoryTaxPreparationServices                                     IssuingCardholderSpendingControlsAllowedCategory = "tax_preparation_services"
	IssuingCardholderSpendingControlsAllowedCategoryTaxicabsLimousines                                         IssuingCardholderSpendingControlsAllowedCategory = "taxicabs_limousines"
	IssuingCardholderSpendingControlsAllowedCategoryTelecommunicationEquipmentAndTelephoneSales                IssuingCardholderSpendingControlsAllowedCategory = "telecommunication_equipment_and_telephone_sales"
	IssuingCardholderSpendingControlsAllowedCategoryTelecommunicationServices                                  IssuingCardholderSpendingControlsAllowedCategory = "telecommunication_services"
	IssuingCardholderSpendingControlsAllowedCategoryTelegraphServices                                          IssuingCardholderSpendingControlsAllowedCategory = "telegraph_services"
	IssuingCardholderSpendingControlsAllowedCategoryTentAndAwningShops                                         IssuingCardholderSpendingControlsAllowedCategory = "tent_and_awning_shops"
	IssuingCardholderSpendingControlsAllowedCategoryTestingLaboratories                                        IssuingCardholderSpendingControlsAllowedCategory = "testing_laboratories"
	IssuingCardholderSpendingControlsAllowedCategoryTheatricalTicketAgencies                                   IssuingCardholderSpendingControlsAllowedCategory = "theatrical_ticket_agencies"
	IssuingCardholderSpendingControlsAllowedCategoryTimeshares                                                 IssuingCardholderSpendingControlsAllowedCategory = "timeshares"
	IssuingCardholderSpendingControlsAllowedCategoryTireRetreadingAndRepair                                    IssuingCardholderSpendingControlsAllowedCategory = "tire_retreading_and_repair"
	IssuingCardholderSpendingControlsAllowedCategoryTollsBridgeFees                                            IssuingCardholderSpendingControlsAllowedCategory = "tolls_bridge_fees"
	IssuingCardholderSpendingControlsAllowedCategoryTouristAttractionsAndExhibits                              IssuingCardholderSpendingControlsAllowedCategory = "tourist_attractions_and_exhibits"
	IssuingCardholderSpendingControlsAllowedCategoryTowingServices                                             IssuingCardholderSpendingControlsAllowedCategory = "towing_services"
	IssuingCardholderSpendingControlsAllowedCategoryTrailerParksCampgrounds                                    IssuingCardholderSpendingControlsAllowedCategory = "trailer_parks_campgrounds"
	IssuingCardholderSpendingControlsAllowedCategoryTransportationServices                                     IssuingCardholderSpendingControlsAllowedCategory = "transportation_services"
	IssuingCardholderSpendingControlsAllowedCategoryTravelAgenciesTourOperators                                IssuingCardholderSpendingControlsAllowedCategory = "travel_agencies_tour_operators"
	IssuingCardholderSpendingControlsAllowedCategoryTruckStopIteration                                         IssuingCardholderSpendingControlsAllowedCategory = "truck_stop_iteration"
	IssuingCardholderSpendingControlsAllowedCategoryTruckUtilityTrailerRentals                                 IssuingCardholderSpendingControlsAllowedCategory = "truck_utility_trailer_rentals"
	IssuingCardholderSpendingControlsAllowedCategoryTypesettingPlateMakingAndRelatedServices                   IssuingCardholderSpendingControlsAllowedCategory = "typesetting_plate_making_and_related_services"
	IssuingCardholderSpendingControlsAllowedCategoryTypewriterStores                                           IssuingCardholderSpendingControlsAllowedCategory = "typewriter_stores"
	IssuingCardholderSpendingControlsAllowedCategoryUSFederalGovernmentAgenciesOrDepartments                   IssuingCardholderSpendingControlsAllowedCategory = "u_s_federal_government_agencies_or_departments"
	IssuingCardholderSpendingControlsAllowedCategoryUniformsCommercialClothing                                 IssuingCardholderSpendingControlsAllowedCategory = "uniforms_commercial_clothing"
	IssuingCardholderSpendingControlsAllowedCategoryUsedMerchandiseAndSecondhandStores                         IssuingCardholderSpendingControlsAllowedCategory = "used_merchandise_and_secondhand_stores"
	IssuingCardholderSpendingControlsAllowedCategoryUtilities                                                  IssuingCardholderSpendingControlsAllowedCategory = "utilities"
	IssuingCardholderSpendingControlsAllowedCategoryVarietyStores                                              IssuingCardholderSpendingControlsAllowedCategory = "variety_stores"
	IssuingCardholderSpendingControlsAllowedCategoryVeterinaryServices                                         IssuingCardholderSpendingControlsAllowedCategory = "veterinary_services"
	IssuingCardholderSpendingControlsAllowedCategoryVideoAmusementGameSupplies                                 IssuingCardholderSpendingControlsAllowedCategory = "video_amusement_game_supplies"
	IssuingCardholderSpendingControlsAllowedCategoryVideoGameArcades                                           IssuingCardholderSpendingControlsAllowedCategory = "video_game_arcades"
	IssuingCardholderSpendingControlsAllowedCategoryVideoTapeRentalStores                                      IssuingCardholderSpendingControlsAllowedCategory = "video_tape_rental_stores"
	IssuingCardholderSpendingControlsAllowedCategoryVocationalTradeSchools                                     IssuingCardholderSpendingControlsAllowedCategory = "vocational_trade_schools"
	IssuingCardholderSpendingControlsAllowedCategoryWatchJewelryRepair                                         IssuingCardholderSpendingControlsAllowedCategory = "watch_jewelry_repair"
	IssuingCardholderSpendingControlsAllowedCategoryWeldingRepair                                              IssuingCardholderSpendingControlsAllowedCategory = "welding_repair"
	IssuingCardholderSpendingControlsAllowedCategoryWholesaleClubs                                             IssuingCardholderSpendingControlsAllowedCategory = "wholesale_clubs"
	IssuingCardholderSpendingControlsAllowedCategoryWigAndToupeeStores                                         IssuingCardholderSpendingControlsAllowedCategory = "wig_and_toupee_stores"
	IssuingCardholderSpendingControlsAllowedCategoryWiresMoneyOrders                                           IssuingCardholderSpendingControlsAllowedCategory = "wires_money_orders"
	IssuingCardholderSpendingControlsAllowedCategoryWomensAccessoryAndSpecialtyShops                           IssuingCardholderSpendingControlsAllowedCategory = "womens_accessory_and_specialty_shops"
	IssuingCardholderSpendingControlsAllowedCategoryWomensReadyToWearStores                                    IssuingCardholderSpendingControlsAllowedCategory = "womens_ready_to_wear_stores"
	IssuingCardholderSpendingControlsAllowedCategoryWreckingAndSalvageYards                                    IssuingCardholderSpendingControlsAllowedCategory = "wrecking_and_salvage_yards"
)

type IssuingCardholderSpendingControlsBlockedCategory string

const (
	IssuingCardholderSpendingControlsBlockedCategoryAcRefrigerationRepair                                      IssuingCardholderSpendingControlsBlockedCategory = "ac_refrigeration_repair"
	IssuingCardholderSpendingControlsBlockedCategoryAccountingBookkeepingServices                              IssuingCardholderSpendingControlsBlockedCategory = "accounting_bookkeeping_services"
	IssuingCardholderSpendingControlsBlockedCategoryAdvertisingServices                                        IssuingCardholderSpendingControlsBlockedCategory = "advertising_services"
	IssuingCardholderSpendingControlsBlockedCategoryAgriculturalCooperative                                    IssuingCardholderSpendingControlsBlockedCategory = "agricultural_cooperative"
	IssuingCardholderSpendingControlsBlockedCategoryAirlinesAirCarriers                                        IssuingCardholderSpendingControlsBlockedCategory = "airlines_air_carriers"
	IssuingCardholderSpendingControlsBlockedCategoryAirportsFlyingFields                                       IssuingCardholderSpendingControlsBlockedCategory = "airports_flying_fields"
	IssuingCardholderSpendingControlsBlockedCategoryAmbulanceServices                                          IssuingCardholderSpendingControlsBlockedCategory = "ambulance_services"
	IssuingCardholderSpendingControlsBlockedCategoryAmusementParksCarnivals                                    IssuingCardholderSpendingControlsBlockedCategory = "amusement_parks_carnivals"
	IssuingCardholderSpendingControlsBlockedCategoryAntiqueReproductions                                       IssuingCardholderSpendingControlsBlockedCategory = "antique_reproductions"
	IssuingCardholderSpendingControlsBlockedCategoryAntiqueShops                                               IssuingCardholderSpendingControlsBlockedCategory = "antique_shops"
	IssuingCardholderSpendingControlsBlockedCategoryAquariums                                                  IssuingCardholderSpendingControlsBlockedCategory = "aquariums"
	IssuingCardholderSpendingControlsBlockedCategoryArchitecturalSurveyingServices                             IssuingCardholderSpendingControlsBlockedCategory = "architectural_surveying_services"
	IssuingCardholderSpendingControlsBlockedCategoryArtDealersAndGalleries                                     IssuingCardholderSpendingControlsBlockedCategory = "art_dealers_and_galleries"
	IssuingCardholderSpendingControlsBlockedCategoryArtistsSupplyAndCraftShops                                 IssuingCardholderSpendingControlsBlockedCategory = "artists_supply_and_craft_shops"
	IssuingCardholderSpendingControlsBlockedCategoryAutoAndHomeSupplyStores                                    IssuingCardholderSpendingControlsBlockedCategory = "auto_and_home_supply_stores"
	IssuingCardholderSpendingControlsBlockedCategoryAutoBodyRepairShops                                        IssuingCardholderSpendingControlsBlockedCategory = "auto_body_repair_shops"
	IssuingCardholderSpendingControlsBlockedCategoryAutoPaintShops                                             IssuingCardholderSpendingControlsBlockedCategory = "auto_paint_shops"
	IssuingCardholderSpendingControlsBlockedCategoryAutoServiceShops                                           IssuingCardholderSpendingControlsBlockedCategory = "auto_service_shops"
	IssuingCardholderSpendingControlsBlockedCategoryAutomatedCashDisburse                                      IssuingCardholderSpendingControlsBlockedCategory = "automated_cash_disburse"
	IssuingCardholderSpendingControlsBlockedCategoryAutomatedFuelDispensers                                    IssuingCardholderSpendingControlsBlockedCategory = "automated_fuel_dispensers"
	IssuingCardholderSpendingControlsBlockedCategoryAutomobileAssociations                                     IssuingCardholderSpendingControlsBlockedCategory = "automobile_associations"
	IssuingCardholderSpendingControlsBlockedCategoryAutomotivePartsAndAccessoriesStores                        IssuingCardholderSpendingControlsBlockedCategory = "automotive_parts_and_accessories_stores"
	IssuingCardholderSpendingControlsBlockedCategoryAutomotiveTireStores                                       IssuingCardholderSpendingControlsBlockedCategory = "automotive_tire_stores"
	IssuingCardholderSpendingControlsBlockedCategoryBailAndBondPayments                                        IssuingCardholderSpendingControlsBlockedCategory = "bail_and_bond_payments"
	IssuingCardholderSpendingControlsBlockedCategoryBakeries                                                   IssuingCardholderSpendingControlsBlockedCategory = "bakeries"
	IssuingCardholderSpendingControlsBlockedCategoryBandsOrchestras                                            IssuingCardholderSpendingControlsBlockedCategory = "bands_orchestras"
	IssuingCardholderSpendingControlsBlockedCategoryBarberAndBeautyShops                                       IssuingCardholderSpendingControlsBlockedCategory = "barber_and_beauty_shops"
	IssuingCardholderSpendingControlsBlockedCategoryBettingCasinoGambling                                      IssuingCardholderSpendingControlsBlockedCategory = "betting_casino_gambling"
	IssuingCardholderSpendingControlsBlockedCategoryBicycleShops                                               IssuingCardholderSpendingControlsBlockedCategory = "bicycle_shops"
	IssuingCardholderSpendingControlsBlockedCategoryBilliardPoolEstablishments                                 IssuingCardholderSpendingControlsBlockedCategory = "billiard_pool_establishments"
	IssuingCardholderSpendingControlsBlockedCategoryBoatDealers                                                IssuingCardholderSpendingControlsBlockedCategory = "boat_dealers"
	IssuingCardholderSpendingControlsBlockedCategoryBoatRentalsAndLeases                                       IssuingCardholderSpendingControlsBlockedCategory = "boat_rentals_and_leases"
	IssuingCardholderSpendingControlsBlockedCategoryBookStores                                                 IssuingCardholderSpendingControlsBlockedCategory = "book_stores"
	IssuingCardholderSpendingControlsBlockedCategoryBooksPeriodicalsAndNewspapers                              IssuingCardholderSpendingControlsBlockedCategory = "books_periodicals_and_newspapers"
	IssuingCardholderSpendingControlsBlockedCategoryBowlingAlleys                                              IssuingCardholderSpendingControlsBlockedCategory = "bowling_alleys"
	IssuingCardholderSpendingControlsBlockedCategoryBusLines                                                   IssuingCardholderSpendingControlsBlockedCategory = "bus_lines"
	IssuingCardholderSpendingControlsBlockedCategoryBusinessSecretarialSchools                                 IssuingCardholderSpendingControlsBlockedCategory = "business_secretarial_schools"
	IssuingCardholderSpendingControlsBlockedCategoryBuyingShoppingServices                                     IssuingCardholderSpendingControlsBlockedCategory = "buying_shopping_services"
	IssuingCardholderSpendingControlsBlockedCategoryCableSatelliteAndOtherPayTelevisionAndRadio                IssuingCardholderSpendingControlsBlockedCategory = "cable_satellite_and_other_pay_television_and_radio"
	IssuingCardholderSpendingControlsBlockedCategoryCameraAndPhotographicSupplyStores                          IssuingCardholderSpendingControlsBlockedCategory = "camera_and_photographic_supply_stores"
	IssuingCardholderSpendingControlsBlockedCategoryCandyNutAndConfectioneryStores                             IssuingCardholderSpendingControlsBlockedCategory = "candy_nut_and_confectionery_stores"
	IssuingCardholderSpendingControlsBlockedCategoryCarAndTruckDealersNewUsed                                  IssuingCardholderSpendingControlsBlockedCategory = "car_and_truck_dealers_new_used"
	IssuingCardholderSpendingControlsBlockedCategoryCarAndTruckDealersUsedOnly                                 IssuingCardholderSpendingControlsBlockedCategory = "car_and_truck_dealers_used_only"
	IssuingCardholderSpendingControlsBlockedCategoryCarRentalAgencies                                          IssuingCardholderSpendingControlsBlockedCategory = "car_rental_agencies"
	IssuingCardholderSpendingControlsBlockedCategoryCarWashes                                                  IssuingCardholderSpendingControlsBlockedCategory = "car_washes"
	IssuingCardholderSpendingControlsBlockedCategoryCarpentryServices                                          IssuingCardholderSpendingControlsBlockedCategory = "carpentry_services"
	IssuingCardholderSpendingControlsBlockedCategoryCarpetUpholsteryCleaning                                   IssuingCardholderSpendingControlsBlockedCategory = "carpet_upholstery_cleaning"
	IssuingCardholderSpendingControlsBlockedCategoryCaterers                                                   IssuingCardholderSpendingControlsBlockedCategory = "caterers"
	IssuingCardholderSpendingControlsBlockedCategoryCharitableAndSocialServiceOrganizationsFundraising         IssuingCardholderSpendingControlsBlockedCategory = "charitable_and_social_service_organizations_fundraising"
	IssuingCardholderSpendingControlsBlockedCategoryChemicalsAndAlliedProducts                                 IssuingCardholderSpendingControlsBlockedCategory = "chemicals_and_allied_products"
	IssuingCardholderSpendingControlsBlockedCategoryChildCareServices                                          IssuingCardholderSpendingControlsBlockedCategory = "child_care_services"
	IssuingCardholderSpendingControlsBlockedCategoryChildrensAndInfantsWearStores                              IssuingCardholderSpendingControlsBlockedCategory = "childrens_and_infants_wear_stores"
	IssuingCardholderSpendingControlsBlockedCategoryChiropodistsPodiatrists                                    IssuingCardholderSpendingControlsBlockedCategory = "chiropodists_podiatrists"
	IssuingCardholderSpendingControlsBlockedCategoryChiropractors                                              IssuingCardholderSpendingControlsBlockedCategory = "chiropractors"
	IssuingCardholderSpendingControlsBlockedCategoryCigarStoresAndStands                                       IssuingCardholderSpendingControlsBlockedCategory = "cigar_stores_and_stands"
	IssuingCardholderSpendingControlsBlockedCategoryCivicSocialFraternalAssociations                           IssuingCardholderSpendingControlsBlockedCategory = "civic_social_fraternal_associations"
	IssuingCardholderSpendingControlsBlockedCategoryCleaningAndMaintenance                                     IssuingCardholderSpendingControlsBlockedCategory = "cleaning_and_maintenance"
	IssuingCardholderSpendingControlsBlockedCategoryClothingRental                                             IssuingCardholderSpendingControlsBlockedCategory = "clothing_rental"
	IssuingCardholderSpendingControlsBlockedCategoryCollegesUniversities                                       IssuingCardholderSpendingControlsBlockedCategory = "colleges_universities"
	IssuingCardholderSpendingControlsBlockedCategoryCommercialEquipment                                        IssuingCardholderSpendingControlsBlockedCategory = "commercial_equipment"
	IssuingCardholderSpendingControlsBlockedCategoryCommercialFootwear                                         IssuingCardholderSpendingControlsBlockedCategory = "commercial_footwear"
	IssuingCardholderSpendingControlsBlockedCategoryCommercialPhotographyArtAndGraphics                        IssuingCardholderSpendingControlsBlockedCategory = "commercial_photography_art_and_graphics"
	IssuingCardholderSpendingControlsBlockedCategoryCommuterTransportAndFerries                                IssuingCardholderSpendingControlsBlockedCategory = "commuter_transport_and_ferries"
	IssuingCardholderSpendingControlsBlockedCategoryComputerNetworkServices                                    IssuingCardholderSpendingControlsBlockedCategory = "computer_network_services"
	IssuingCardholderSpendingControlsBlockedCategoryComputerProgramming                                        IssuingCardholderSpendingControlsBlockedCategory = "computer_programming"
	IssuingCardholderSpendingControlsBlockedCategoryComputerRepair                                             IssuingCardholderSpendingControlsBlockedCategory = "computer_repair"
	IssuingCardholderSpendingControlsBlockedCategoryComputerSoftwareStores                                     IssuingCardholderSpendingControlsBlockedCategory = "computer_software_stores"
	IssuingCardholderSpendingControlsBlockedCategoryComputersPeripheralsAndSoftware                            IssuingCardholderSpendingControlsBlockedCategory = "computers_peripherals_and_software"
	IssuingCardholderSpendingControlsBlockedCategoryConcreteWorkServices                                       IssuingCardholderSpendingControlsBlockedCategory = "concrete_work_services"
	IssuingCardholderSpendingControlsBlockedCategoryConstructionMaterials                                      IssuingCardholderSpendingControlsBlockedCategory = "construction_materials"
	IssuingCardholderSpendingControlsBlockedCategoryConsultingPublicRelations                                  IssuingCardholderSpendingControlsBlockedCategory = "consulting_public_relations"
	IssuingCardholderSpendingControlsBlockedCategoryCorrespondenceSchools                                      IssuingCardholderSpendingControlsBlockedCategory = "correspondence_schools"
	IssuingCardholderSpendingControlsBlockedCategoryCosmeticStores                                             IssuingCardholderSpendingControlsBlockedCategory = "cosmetic_stores"
	IssuingCardholderSpendingControlsBlockedCategoryCounselingServices                                         IssuingCardholderSpendingControlsBlockedCategory = "counseling_services"
	IssuingCardholderSpendingControlsBlockedCategoryCountryClubs                                               IssuingCardholderSpendingControlsBlockedCategory = "country_clubs"
	IssuingCardholderSpendingControlsBlockedCategoryCourierServices                                            IssuingCardholderSpendingControlsBlockedCategory = "courier_services"
	IssuingCardholderSpendingControlsBlockedCategoryCourtCosts                                                 IssuingCardholderSpendingControlsBlockedCategory = "court_costs"
	IssuingCardholderSpendingControlsBlockedCategoryCreditReportingAgencies                                    IssuingCardholderSpendingControlsBlockedCategory = "credit_reporting_agencies"
	IssuingCardholderSpendingControlsBlockedCategoryCruiseLines                                                IssuingCardholderSpendingControlsBlockedCategory = "cruise_lines"
	IssuingCardholderSpendingControlsBlockedCategoryDairyProductsStores                                        IssuingCardholderSpendingControlsBlockedCategory = "dairy_products_stores"
	IssuingCardholderSpendingControlsBlockedCategoryDanceHallStudiosSchools                                    IssuingCardholderSpendingControlsBlockedCategory = "dance_hall_studios_schools"
	IssuingCardholderSpendingControlsBlockedCategoryDatingEscortServices                                       IssuingCardholderSpendingControlsBlockedCategory = "dating_escort_services"
	IssuingCardholderSpendingControlsBlockedCategoryDentistsOrthodontists                                      IssuingCardholderSpendingControlsBlockedCategory = "dentists_orthodontists"
	IssuingCardholderSpendingControlsBlockedCategoryDepartmentStores                                           IssuingCardholderSpendingControlsBlockedCategory = "department_stores"
	IssuingCardholderSpendingControlsBlockedCategoryDetectiveAgencies                                          IssuingCardholderSpendingControlsBlockedCategory = "detective_agencies"
	IssuingCardholderSpendingControlsBlockedCategoryDigitalGoodsApplications                                   IssuingCardholderSpendingControlsBlockedCategory = "digital_goods_applications"
	IssuingCardholderSpendingControlsBlockedCategoryDigitalGoodsGames                                          IssuingCardholderSpendingControlsBlockedCategory = "digital_goods_games"
	IssuingCardholderSpendingControlsBlockedCategoryDigitalGoodsLargeVolume                                    IssuingCardholderSpendingControlsBlockedCategory = "digital_goods_large_volume"
	IssuingCardholderSpendingControlsBlockedCategoryDigitalGoodsMedia                                          IssuingCardholderSpendingControlsBlockedCategory = "digital_goods_media"
	IssuingCardholderSpendingControlsBlockedCategoryDirectMarketingCatalogMerchant                             IssuingCardholderSpendingControlsBlockedCategory = "direct_marketing_catalog_merchant"
	IssuingCardholderSpendingControlsBlockedCategoryDirectMarketingCombinationCatalogAndRetailMerchant         IssuingCardholderSpendingControlsBlockedCategory = "direct_marketing_combination_catalog_and_retail_merchant"
	IssuingCardholderSpendingControlsBlockedCategoryDirectMarketingInboundTelemarketing                        IssuingCardholderSpendingControlsBlockedCategory = "direct_marketing_inbound_telemarketing"
	IssuingCardholderSpendingControlsBlockedCategoryDirectMarketingInsuranceServices                           IssuingCardholderSpendingControlsBlockedCategory = "direct_marketing_insurance_services"
	IssuingCardholderSpendingControlsBlockedCategoryDirectMarketingOther                                       IssuingCardholderSpendingControlsBlockedCategory = "direct_marketing_other"
	IssuingCardholderSpendingControlsBlockedCategoryDirectMarketingOutboundTelemarketing                       IssuingCardholderSpendingControlsBlockedCategory = "direct_marketing_outbound_telemarketing"
	IssuingCardholderSpendingControlsBlockedCategoryDirectMarketingSubscription                                IssuingCardholderSpendingControlsBlockedCategory = "direct_marketing_subscription"
	IssuingCardholderSpendingControlsBlockedCategoryDirectMarketingTravel                                      IssuingCardholderSpendingControlsBlockedCategory = "direct_marketing_travel"
	IssuingCardholderSpendingControlsBlockedCategoryDiscountStores                                             IssuingCardholderSpendingControlsBlockedCategory = "discount_stores"
	IssuingCardholderSpendingControlsBlockedCategoryDoctors                                                    IssuingCardholderSpendingControlsBlockedCategory = "doctors"
	IssuingCardholderSpendingControlsBlockedCategoryDoorToDoorSales                                            IssuingCardholderSpendingControlsBlockedCategory = "door_to_door_sales"
	IssuingCardholderSpendingControlsBlockedCategoryDraperyWindowCoveringAndUpholsteryStores                   IssuingCardholderSpendingControlsBlockedCategory = "drapery_window_covering_and_upholstery_stores"
	IssuingCardholderSpendingControlsBlockedCategoryDrinkingPlaces                                             IssuingCardholderSpendingControlsBlockedCategory = "drinking_places"
	IssuingCardholderSpendingControlsBlockedCategoryDrugStoresAndPharmacies                                    IssuingCardholderSpendingControlsBlockedCategory = "drug_stores_and_pharmacies"
	IssuingCardholderSpendingControlsBlockedCategoryDrugsDrugProprietariesAndDruggistSundries                  IssuingCardholderSpendingControlsBlockedCategory = "drugs_drug_proprietaries_and_druggist_sundries"
	IssuingCardholderSpendingControlsBlockedCategoryDryCleaners                                                IssuingCardholderSpendingControlsBlockedCategory = "dry_cleaners"
	IssuingCardholderSpendingControlsBlockedCategoryDurableGoods                                               IssuingCardholderSpendingControlsBlockedCategory = "durable_goods"
	IssuingCardholderSpendingControlsBlockedCategoryDutyFreeStores                                             IssuingCardholderSpendingControlsBlockedCategory = "duty_free_stores"
	IssuingCardholderSpendingControlsBlockedCategoryEatingPlacesRestaurants                                    IssuingCardholderSpendingControlsBlockedCategory = "eating_places_restaurants"
	IssuingCardholderSpendingControlsBlockedCategoryEducationalServices                                        IssuingCardholderSpendingControlsBlockedCategory = "educational_services"
	IssuingCardholderSpendingControlsBlockedCategoryElectricRazorStores                                        IssuingCardholderSpendingControlsBlockedCategory = "electric_razor_stores"
	IssuingCardholderSpendingControlsBlockedCategoryElectricalPartsAndEquipment                                IssuingCardholderSpendingControlsBlockedCategory = "electrical_parts_and_equipment"
	IssuingCardholderSpendingControlsBlockedCategoryElectricalServices                                         IssuingCardholderSpendingControlsBlockedCategory = "electrical_services"
	IssuingCardholderSpendingControlsBlockedCategoryElectronicsRepairShops                                     IssuingCardholderSpendingControlsBlockedCategory = "electronics_repair_shops"
	IssuingCardholderSpendingControlsBlockedCategoryElectronicsStores                                          IssuingCardholderSpendingControlsBlockedCategory = "electronics_stores"
	IssuingCardholderSpendingControlsBlockedCategoryElementarySecondarySchools                                 IssuingCardholderSpendingControlsBlockedCategory = "elementary_secondary_schools"
	IssuingCardholderSpendingControlsBlockedCategoryEmploymentTempAgencies                                     IssuingCardholderSpendingControlsBlockedCategory = "employment_temp_agencies"
	IssuingCardholderSpendingControlsBlockedCategoryEquipmentRental                                            IssuingCardholderSpendingControlsBlockedCategory = "equipment_rental"
	IssuingCardholderSpendingControlsBlockedCategoryExterminatingServices                                      IssuingCardholderSpendingControlsBlockedCategory = "exterminating_services"
	IssuingCardholderSpendingControlsBlockedCategoryFamilyClothingStores                                       IssuingCardholderSpendingControlsBlockedCategory = "family_clothing_stores"
	IssuingCardholderSpendingControlsBlockedCategoryFastFoodRestaurants                                        IssuingCardholderSpendingControlsBlockedCategory = "fast_food_restaurants"
	IssuingCardholderSpendingControlsBlockedCategoryFinancialInstitutions                                      IssuingCardholderSpendingControlsBlockedCategory = "financial_institutions"
	IssuingCardholderSpendingControlsBlockedCategoryFinesGovernmentAdministrativeEntities                      IssuingCardholderSpendingControlsBlockedCategory = "fines_government_administrative_entities"
	IssuingCardholderSpendingControlsBlockedCategoryFireplaceFireplaceScreensAndAccessoriesStores              IssuingCardholderSpendingControlsBlockedCategory = "fireplace_fireplace_screens_and_accessories_stores"
	IssuingCardholderSpendingControlsBlockedCategoryFloorCoveringStores                                        IssuingCardholderSpendingControlsBlockedCategory = "floor_covering_stores"
	IssuingCardholderSpendingControlsBlockedCategoryFlorists                                                   IssuingCardholderSpendingControlsBlockedCategory = "florists"
	IssuingCardholderSpendingControlsBlockedCategoryFloristsSuppliesNurseryStockAndFlowers                     IssuingCardholderSpendingControlsBlockedCategory = "florists_supplies_nursery_stock_and_flowers"
	IssuingCardholderSpendingControlsBlockedCategoryFreezerAndLockerMeatProvisioners                           IssuingCardholderSpendingControlsBlockedCategory = "freezer_and_locker_meat_provisioners"
	IssuingCardholderSpendingControlsBlockedCategoryFuelDealersNonAutomotive                                   IssuingCardholderSpendingControlsBlockedCategory = "fuel_dealers_non_automotive"
	IssuingCardholderSpendingControlsBlockedCategoryFuneralServicesCrematories                                 IssuingCardholderSpendingControlsBlockedCategory = "funeral_services_crematories"
	IssuingCardholderSpendingControlsBlockedCategoryFurnitureHomeFurnishingsAndEquipmentStoresExceptAppliances IssuingCardholderSpendingControlsBlockedCategory = "furniture_home_furnishings_and_equipment_stores_except_appliances"
	IssuingCardholderSpendingControlsBlockedCategoryFurnitureRepairRefinishing                                 IssuingCardholderSpendingControlsBlockedCategory = "furniture_repair_refinishing"
	IssuingCardholderSpendingControlsBlockedCategoryFurriersAndFurShops                                        IssuingCardholderSpendingControlsBlockedCategory = "furriers_and_fur_shops"
	IssuingCardholderSpendingControlsBlockedCategoryGeneralServices                                            IssuingCardholderSpendingControlsBlockedCategory = "general_services"
	IssuingCardholderSpendingControlsBlockedCategoryGiftCardNoveltyAndSouvenirShops                            IssuingCardholderSpendingControlsBlockedCategory = "gift_card_novelty_and_souvenir_shops"
	IssuingCardholderSpendingControlsBlockedCategoryGlassPaintAndWallpaperStores                               IssuingCardholderSpendingControlsBlockedCategory = "glass_paint_and_wallpaper_stores"
	IssuingCardholderSpendingControlsBlockedCategoryGlasswareCrystalStores                                     IssuingCardholderSpendingControlsBlockedCategory = "glassware_crystal_stores"
	IssuingCardholderSpendingControlsBlockedCategoryGolfCoursesPublic                                          IssuingCardholderSpendingControlsBlockedCategory = "golf_courses_public"
	IssuingCardholderSpendingControlsBlockedCategoryGovernmentServices                                         IssuingCardholderSpendingControlsBlockedCategory = "government_services"
	IssuingCardholderSpendingControlsBlockedCategoryGroceryStoresSupermarkets                                  IssuingCardholderSpendingControlsBlockedCategory = "grocery_stores_supermarkets"
	IssuingCardholderSpendingControlsBlockedCategoryHardwareEquipmentAndSupplies                               IssuingCardholderSpendingControlsBlockedCategory = "hardware_equipment_and_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryHardwareStores                                             IssuingCardholderSpendingControlsBlockedCategory = "hardware_stores"
	IssuingCardholderSpendingControlsBlockedCategoryHealthAndBeautySpas                                        IssuingCardholderSpendingControlsBlockedCategory = "health_and_beauty_spas"
	IssuingCardholderSpendingControlsBlockedCategoryHearingAidsSalesAndSupplies                                IssuingCardholderSpendingControlsBlockedCategory = "hearing_aids_sales_and_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryHeatingPlumbingAC                                          IssuingCardholderSpendingControlsBlockedCategory = "heating_plumbing_a_c"
	IssuingCardholderSpendingControlsBlockedCategoryHobbyToyAndGameShops                                       IssuingCardholderSpendingControlsBlockedCategory = "hobby_toy_and_game_shops"
	IssuingCardholderSpendingControlsBlockedCategoryHomeSupplyWarehouseStores                                  IssuingCardholderSpendingControlsBlockedCategory = "home_supply_warehouse_stores"
	IssuingCardholderSpendingControlsBlockedCategoryHospitals                                                  IssuingCardholderSpendingControlsBlockedCategory = "hospitals"
	IssuingCardholderSpendingControlsBlockedCategoryHotelsMotelsAndResorts                                     IssuingCardholderSpendingControlsBlockedCategory = "hotels_motels_and_resorts"
	IssuingCardholderSpendingControlsBlockedCategoryHouseholdApplianceStores                                   IssuingCardholderSpendingControlsBlockedCategory = "household_appliance_stores"
	IssuingCardholderSpendingControlsBlockedCategoryIndustrialSupplies                                         IssuingCardholderSpendingControlsBlockedCategory = "industrial_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryInformationRetrievalServices                               IssuingCardholderSpendingControlsBlockedCategory = "information_retrieval_services"
	IssuingCardholderSpendingControlsBlockedCategoryInsuranceDefault                                           IssuingCardholderSpendingControlsBlockedCategory = "insurance_default"
	IssuingCardholderSpendingControlsBlockedCategoryInsuranceUnderwritingPremiums                              IssuingCardholderSpendingControlsBlockedCategory = "insurance_underwriting_premiums"
	IssuingCardholderSpendingControlsBlockedCategoryIntraCompanyPurchases                                      IssuingCardholderSpendingControlsBlockedCategory = "intra_company_purchases"
	IssuingCardholderSpendingControlsBlockedCategoryJewelryStoresWatchesClocksAndSilverwareStores              IssuingCardholderSpendingControlsBlockedCategory = "jewelry_stores_watches_clocks_and_silverware_stores"
	IssuingCardholderSpendingControlsBlockedCategoryLandscapingServices                                        IssuingCardholderSpendingControlsBlockedCategory = "landscaping_services"
	IssuingCardholderSpendingControlsBlockedCategoryLaundries                                                  IssuingCardholderSpendingControlsBlockedCategory = "laundries"
	IssuingCardholderSpendingControlsBlockedCategoryLaundryCleaningServices                                    IssuingCardholderSpendingControlsBlockedCategory = "laundry_cleaning_services"
	IssuingCardholderSpendingControlsBlockedCategoryLegalServicesAttorneys                                     IssuingCardholderSpendingControlsBlockedCategory = "legal_services_attorneys"
	IssuingCardholderSpendingControlsBlockedCategoryLuggageAndLeatherGoodsStores                               IssuingCardholderSpendingControlsBlockedCategory = "luggage_and_leather_goods_stores"
	IssuingCardholderSpendingControlsBlockedCategoryLumberBuildingMaterialsStores                              IssuingCardholderSpendingControlsBlockedCategory = "lumber_building_materials_stores"
	IssuingCardholderSpendingControlsBlockedCategoryManualCashDisburse                                         IssuingCardholderSpendingControlsBlockedCategory = "manual_cash_disburse"
	IssuingCardholderSpendingControlsBlockedCategoryMarinasServiceAndSupplies                                  IssuingCardholderSpendingControlsBlockedCategory = "marinas_service_and_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryMasonryStoneworkAndPlaster                                 IssuingCardholderSpendingControlsBlockedCategory = "masonry_stonework_and_plaster"
	IssuingCardholderSpendingControlsBlockedCategoryMassageParlors                                             IssuingCardholderSpendingControlsBlockedCategory = "massage_parlors"
	IssuingCardholderSpendingControlsBlockedCategoryMedicalAndDentalLabs                                       IssuingCardholderSpendingControlsBlockedCategory = "medical_and_dental_labs"
	IssuingCardholderSpendingControlsBlockedCategoryMedicalDentalOphthalmicAndHospitalEquipmentAndSupplies     IssuingCardholderSpendingControlsBlockedCategory = "medical_dental_ophthalmic_and_hospital_equipment_and_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryMedicalServices                                            IssuingCardholderSpendingControlsBlockedCategory = "medical_services"
	IssuingCardholderSpendingControlsBlockedCategoryMembershipOrganizations                                    IssuingCardholderSpendingControlsBlockedCategory = "membership_organizations"
	IssuingCardholderSpendingControlsBlockedCategoryMensAndBoysClothingAndAccessoriesStores                    IssuingCardholderSpendingControlsBlockedCategory = "mens_and_boys_clothing_and_accessories_stores"
	IssuingCardholderSpendingControlsBlockedCategoryMensWomensClothingStores                                   IssuingCardholderSpendingControlsBlockedCategory = "mens_womens_clothing_stores"
	IssuingCardholderSpendingControlsBlockedCategoryMetalServiceCenters                                        IssuingCardholderSpendingControlsBlockedCategory = "metal_service_centers"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneous                                              IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousApparelAndAccessoryShops                      IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_apparel_and_accessory_shops"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousAutoDealers                                   IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_auto_dealers"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousBusinessServices                              IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_business_services"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousFoodStores                                    IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_food_stores"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousGeneralMerchandise                            IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_general_merchandise"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousGeneralServices                               IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_general_services"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousHomeFurnishingSpecialtyStores                 IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_home_furnishing_specialty_stores"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousPublishingAndPrinting                         IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_publishing_and_printing"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousRecreationServices                            IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_recreation_services"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousRepairShops                                   IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_repair_shops"
	IssuingCardholderSpendingControlsBlockedCategoryMiscellaneousSpecialtyRetail                               IssuingCardholderSpendingControlsBlockedCategory = "miscellaneous_specialty_retail"
	IssuingCardholderSpendingControlsBlockedCategoryMobileHomeDealers                                          IssuingCardholderSpendingControlsBlockedCategory = "mobile_home_dealers"
	IssuingCardholderSpendingControlsBlockedCategoryMotionPictureTheaters                                      IssuingCardholderSpendingControlsBlockedCategory = "motion_picture_theaters"
	IssuingCardholderSpendingControlsBlockedCategoryMotorFreightCarriersAndTrucking                            IssuingCardholderSpendingControlsBlockedCategory = "motor_freight_carriers_and_trucking"
	IssuingCardholderSpendingControlsBlockedCategoryMotorHomesDealers                                          IssuingCardholderSpendingControlsBlockedCategory = "motor_homes_dealers"
	IssuingCardholderSpendingControlsBlockedCategoryMotorVehicleSuppliesAndNewParts                            IssuingCardholderSpendingControlsBlockedCategory = "motor_vehicle_supplies_and_new_parts"
	IssuingCardholderSpendingControlsBlockedCategoryMotorcycleShopsAndDealers                                  IssuingCardholderSpendingControlsBlockedCategory = "motorcycle_shops_and_dealers"
	IssuingCardholderSpendingControlsBlockedCategoryMotorcycleShopsDealers                                     IssuingCardholderSpendingControlsBlockedCategory = "motorcycle_shops_dealers"
	IssuingCardholderSpendingControlsBlockedCategoryMusicStoresMusicalInstrumentsPianosAndSheetMusic           IssuingCardholderSpendingControlsBlockedCategory = "music_stores_musical_instruments_pianos_and_sheet_music"
	IssuingCardholderSpendingControlsBlockedCategoryNewsDealersAndNewsstands                                   IssuingCardholderSpendingControlsBlockedCategory = "news_dealers_and_newsstands"
	IssuingCardholderSpendingControlsBlockedCategoryNonFiMoneyOrders                                           IssuingCardholderSpendingControlsBlockedCategory = "non_fi_money_orders"
	IssuingCardholderSpendingControlsBlockedCategoryNonFiStoredValueCardPurchaseLoad                           IssuingCardholderSpendingControlsBlockedCategory = "non_fi_stored_value_card_purchase_load"
	IssuingCardholderSpendingControlsBlockedCategoryNondurableGoods                                            IssuingCardholderSpendingControlsBlockedCategory = "nondurable_goods"
	IssuingCardholderSpendingControlsBlockedCategoryNurseriesLawnAndGardenSupplyStores                         IssuingCardholderSpendingControlsBlockedCategory = "nurseries_lawn_and_garden_supply_stores"
	IssuingCardholderSpendingControlsBlockedCategoryNursingPersonalCare                                        IssuingCardholderSpendingControlsBlockedCategory = "nursing_personal_care"
	IssuingCardholderSpendingControlsBlockedCategoryOfficeAndCommercialFurniture                               IssuingCardholderSpendingControlsBlockedCategory = "office_and_commercial_furniture"
	IssuingCardholderSpendingControlsBlockedCategoryOpticiansEyeglasses                                        IssuingCardholderSpendingControlsBlockedCategory = "opticians_eyeglasses"
	IssuingCardholderSpendingControlsBlockedCategoryOptometristsOphthalmologist                                IssuingCardholderSpendingControlsBlockedCategory = "optometrists_ophthalmologist"
	IssuingCardholderSpendingControlsBlockedCategoryOrthopedicGoodsProstheticDevices                           IssuingCardholderSpendingControlsBlockedCategory = "orthopedic_goods_prosthetic_devices"
	IssuingCardholderSpendingControlsBlockedCategoryOsteopaths                                                 IssuingCardholderSpendingControlsBlockedCategory = "osteopaths"
	IssuingCardholderSpendingControlsBlockedCategoryPackageStoresBeerWineAndLiquor                             IssuingCardholderSpendingControlsBlockedCategory = "package_stores_beer_wine_and_liquor"
	IssuingCardholderSpendingControlsBlockedCategoryPaintsVarnishesAndSupplies                                 IssuingCardholderSpendingControlsBlockedCategory = "paints_varnishes_and_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryParkingLotsGarages                                         IssuingCardholderSpendingControlsBlockedCategory = "parking_lots_garages"
	IssuingCardholderSpendingControlsBlockedCategoryPassengerRailways                                          IssuingCardholderSpendingControlsBlockedCategory = "passenger_railways"
	IssuingCardholderSpendingControlsBlockedCategoryPawnShops                                                  IssuingCardholderSpendingControlsBlockedCategory = "pawn_shops"
	IssuingCardholderSpendingControlsBlockedCategoryPetShopsPetFoodAndSupplies                                 IssuingCardholderSpendingControlsBlockedCategory = "pet_shops_pet_food_and_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryPetroleumAndPetroleumProducts                              IssuingCardholderSpendingControlsBlockedCategory = "petroleum_and_petroleum_products"
	IssuingCardholderSpendingControlsBlockedCategoryPhotoDeveloping                                            IssuingCardholderSpendingControlsBlockedCategory = "photo_developing"
	IssuingCardholderSpendingControlsBlockedCategoryPhotographicPhotocopyMicrofilmEquipmentAndSupplies         IssuingCardholderSpendingControlsBlockedCategory = "photographic_photocopy_microfilm_equipment_and_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryPhotographicStudios                                        IssuingCardholderSpendingControlsBlockedCategory = "photographic_studios"
	IssuingCardholderSpendingControlsBlockedCategoryPictureVideoProduction                                     IssuingCardholderSpendingControlsBlockedCategory = "picture_video_production"
	IssuingCardholderSpendingControlsBlockedCategoryPieceGoodsNotionsAndOtherDryGoods                          IssuingCardholderSpendingControlsBlockedCategory = "piece_goods_notions_and_other_dry_goods"
	IssuingCardholderSpendingControlsBlockedCategoryPlumbingHeatingEquipmentAndSupplies                        IssuingCardholderSpendingControlsBlockedCategory = "plumbing_heating_equipment_and_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryPoliticalOrganizations                                     IssuingCardholderSpendingControlsBlockedCategory = "political_organizations"
	IssuingCardholderSpendingControlsBlockedCategoryPostalServicesGovernmentOnly                               IssuingCardholderSpendingControlsBlockedCategory = "postal_services_government_only"
	IssuingCardholderSpendingControlsBlockedCategoryPreciousStonesAndMetalsWatchesAndJewelry                   IssuingCardholderSpendingControlsBlockedCategory = "precious_stones_and_metals_watches_and_jewelry"
	IssuingCardholderSpendingControlsBlockedCategoryProfessionalServices                                       IssuingCardholderSpendingControlsBlockedCategory = "professional_services"
	IssuingCardholderSpendingControlsBlockedCategoryPublicWarehousingAndStorage                                IssuingCardholderSpendingControlsBlockedCategory = "public_warehousing_and_storage"
	IssuingCardholderSpendingControlsBlockedCategoryQuickCopyReproAndBlueprint                                 IssuingCardholderSpendingControlsBlockedCategory = "quick_copy_repro_and_blueprint"
	IssuingCardholderSpendingControlsBlockedCategoryRailroads                                                  IssuingCardholderSpendingControlsBlockedCategory = "railroads"
	IssuingCardholderSpendingControlsBlockedCategoryRealEstateAgentsAndManagersRentals                         IssuingCardholderSpendingControlsBlockedCategory = "real_estate_agents_and_managers_rentals"
	IssuingCardholderSpendingControlsBlockedCategoryRecordStores                                               IssuingCardholderSpendingControlsBlockedCategory = "record_stores"
	IssuingCardholderSpendingControlsBlockedCategoryRecreationalVehicleRentals                                 IssuingCardholderSpendingControlsBlockedCategory = "recreational_vehicle_rentals"
	IssuingCardholderSpendingControlsBlockedCategoryReligiousGoodsStores                                       IssuingCardholderSpendingControlsBlockedCategory = "religious_goods_stores"
	IssuingCardholderSpendingControlsBlockedCategoryReligiousOrganizations                                     IssuingCardholderSpendingControlsBlockedCategory = "religious_organizations"
	IssuingCardholderSpendingControlsBlockedCategoryRoofingSidingSheetMetal                                    IssuingCardholderSpendingControlsBlockedCategory = "roofing_siding_sheet_metal"
	IssuingCardholderSpendingControlsBlockedCategorySecretarialSupportServices                                 IssuingCardholderSpendingControlsBlockedCategory = "secretarial_support_services"
	IssuingCardholderSpendingControlsBlockedCategorySecurityBrokersDealers                                     IssuingCardholderSpendingControlsBlockedCategory = "security_brokers_dealers"
	IssuingCardholderSpendingControlsBlockedCategoryServiceStations                                            IssuingCardholderSpendingControlsBlockedCategory = "service_stations"
	IssuingCardholderSpendingControlsBlockedCategorySewingNeedleworkFabricAndPieceGoodsStores                  IssuingCardholderSpendingControlsBlockedCategory = "sewing_needlework_fabric_and_piece_goods_stores"
	IssuingCardholderSpendingControlsBlockedCategoryShoeRepairHatCleaning                                      IssuingCardholderSpendingControlsBlockedCategory = "shoe_repair_hat_cleaning"
	IssuingCardholderSpendingControlsBlockedCategoryShoeStores                                                 IssuingCardholderSpendingControlsBlockedCategory = "shoe_stores"
	IssuingCardholderSpendingControlsBlockedCategorySmallApplianceRepair                                       IssuingCardholderSpendingControlsBlockedCategory = "small_appliance_repair"
	IssuingCardholderSpendingControlsBlockedCategorySnowmobileDealers                                          IssuingCardholderSpendingControlsBlockedCategory = "snowmobile_dealers"
	IssuingCardholderSpendingControlsBlockedCategorySpecialTradeServices                                       IssuingCardholderSpendingControlsBlockedCategory = "special_trade_services"
	IssuingCardholderSpendingControlsBlockedCategorySpecialtyCleaning                                          IssuingCardholderSpendingControlsBlockedCategory = "specialty_cleaning"
	IssuingCardholderSpendingControlsBlockedCategorySportingGoodsStores                                        IssuingCardholderSpendingControlsBlockedCategory = "sporting_goods_stores"
	IssuingCardholderSpendingControlsBlockedCategorySportingRecreationCamps                                    IssuingCardholderSpendingControlsBlockedCategory = "sporting_recreation_camps"
	IssuingCardholderSpendingControlsBlockedCategorySportsAndRidingApparelStores                               IssuingCardholderSpendingControlsBlockedCategory = "sports_and_riding_apparel_stores"
	IssuingCardholderSpendingControlsBlockedCategorySportsClubsFields                                          IssuingCardholderSpendingControlsBlockedCategory = "sports_clubs_fields"
	IssuingCardholderSpendingControlsBlockedCategoryStampAndCoinStores                                         IssuingCardholderSpendingControlsBlockedCategory = "stamp_and_coin_stores"
	IssuingCardholderSpendingControlsBlockedCategoryStationaryOfficeSuppliesPrintingAndWritingPaper            IssuingCardholderSpendingControlsBlockedCategory = "stationary_office_supplies_printing_and_writing_paper"
	IssuingCardholderSpendingControlsBlockedCategoryStationeryStoresOfficeAndSchoolSupplyStores                IssuingCardholderSpendingControlsBlockedCategory = "stationery_stores_office_and_school_supply_stores"
	IssuingCardholderSpendingControlsBlockedCategorySwimmingPoolsSales                                         IssuingCardholderSpendingControlsBlockedCategory = "swimming_pools_sales"
	IssuingCardholderSpendingControlsBlockedCategoryTUiTravelGermany                                           IssuingCardholderSpendingControlsBlockedCategory = "t_ui_travel_germany"
	IssuingCardholderSpendingControlsBlockedCategoryTailorsAlterations                                         IssuingCardholderSpendingControlsBlockedCategory = "tailors_alterations"
	IssuingCardholderSpendingControlsBlockedCategoryTaxPaymentsGovernmentAgencies                              IssuingCardholderSpendingControlsBlockedCategory = "tax_payments_government_agencies"
	IssuingCardholderSpendingControlsBlockedCategoryTaxPreparationServices                                     IssuingCardholderSpendingControlsBlockedCategory = "tax_preparation_services"
	IssuingCardholderSpendingControlsBlockedCategoryTaxicabsLimousines                                         IssuingCardholderSpendingControlsBlockedCategory = "taxicabs_limousines"
	IssuingCardholderSpendingControlsBlockedCategoryTelecommunicationEquipmentAndTelephoneSales                IssuingCardholderSpendingControlsBlockedCategory = "telecommunication_equipment_and_telephone_sales"
	IssuingCardholderSpendingControlsBlockedCategoryTelecommunicationServices                                  IssuingCardholderSpendingControlsBlockedCategory = "telecommunication_services"
	IssuingCardholderSpendingControlsBlockedCategoryTelegraphServices                                          IssuingCardholderSpendingControlsBlockedCategory = "telegraph_services"
	IssuingCardholderSpendingControlsBlockedCategoryTentAndAwningShops                                         IssuingCardholderSpendingControlsBlockedCategory = "tent_and_awning_shops"
	IssuingCardholderSpendingControlsBlockedCategoryTestingLaboratories                                        IssuingCardholderSpendingControlsBlockedCategory = "testing_laboratories"
	IssuingCardholderSpendingControlsBlockedCategoryTheatricalTicketAgencies                                   IssuingCardholderSpendingControlsBlockedCategory = "theatrical_ticket_agencies"
	IssuingCardholderSpendingControlsBlockedCategoryTimeshares                                                 IssuingCardholderSpendingControlsBlockedCategory = "timeshares"
	IssuingCardholderSpendingControlsBlockedCategoryTireRetreadingAndRepair                                    IssuingCardholderSpendingControlsBlockedCategory = "tire_retreading_and_repair"
	IssuingCardholderSpendingControlsBlockedCategoryTollsBridgeFees                                            IssuingCardholderSpendingControlsBlockedCategory = "tolls_bridge_fees"
	IssuingCardholderSpendingControlsBlockedCategoryTouristAttractionsAndExhibits                              IssuingCardholderSpendingControlsBlockedCategory = "tourist_attractions_and_exhibits"
	IssuingCardholderSpendingControlsBlockedCategoryTowingServices                                             IssuingCardholderSpendingControlsBlockedCategory = "towing_services"
	IssuingCardholderSpendingControlsBlockedCategoryTrailerParksCampgrounds                                    IssuingCardholderSpendingControlsBlockedCategory = "trailer_parks_campgrounds"
	IssuingCardholderSpendingControlsBlockedCategoryTransportationServices                                     IssuingCardholderSpendingControlsBlockedCategory = "transportation_services"
	IssuingCardholderSpendingControlsBlockedCategoryTravelAgenciesTourOperators                                IssuingCardholderSpendingControlsBlockedCategory = "travel_agencies_tour_operators"
	IssuingCardholderSpendingControlsBlockedCategoryTruckStopIteration                                         IssuingCardholderSpendingControlsBlockedCategory = "truck_stop_iteration"
	IssuingCardholderSpendingControlsBlockedCategoryTruckUtilityTrailerRentals                                 IssuingCardholderSpendingControlsBlockedCategory = "truck_utility_trailer_rentals"
	IssuingCardholderSpendingControlsBlockedCategoryTypesettingPlateMakingAndRelatedServices                   IssuingCardholderSpendingControlsBlockedCategory = "typesetting_plate_making_and_related_services"
	IssuingCardholderSpendingControlsBlockedCategoryTypewriterStores                                           IssuingCardholderSpendingControlsBlockedCategory = "typewriter_stores"
	IssuingCardholderSpendingControlsBlockedCategoryUSFederalGovernmentAgenciesOrDepartments                   IssuingCardholderSpendingControlsBlockedCategory = "u_s_federal_government_agencies_or_departments"
	IssuingCardholderSpendingControlsBlockedCategoryUniformsCommercialClothing                                 IssuingCardholderSpendingControlsBlockedCategory = "uniforms_commercial_clothing"
	IssuingCardholderSpendingControlsBlockedCategoryUsedMerchandiseAndSecondhandStores                         IssuingCardholderSpendingControlsBlockedCategory = "used_merchandise_and_secondhand_stores"
	IssuingCardholderSpendingControlsBlockedCategoryUtilities                                                  IssuingCardholderSpendingControlsBlockedCategory = "utilities"
	IssuingCardholderSpendingControlsBlockedCategoryVarietyStores                                              IssuingCardholderSpendingControlsBlockedCategory = "variety_stores"
	IssuingCardholderSpendingControlsBlockedCategoryVeterinaryServices                                         IssuingCardholderSpendingControlsBlockedCategory = "veterinary_services"
	IssuingCardholderSpendingControlsBlockedCategoryVideoAmusementGameSupplies                                 IssuingCardholderSpendingControlsBlockedCategory = "video_amusement_game_supplies"
	IssuingCardholderSpendingControlsBlockedCategoryVideoGameArcades                                           IssuingCardholderSpendingControlsBlockedCategory = "video_game_arcades"
	IssuingCardholderSpendingControlsBlockedCategoryVideoTapeRentalStores                                      IssuingCardholderSpendingControlsBlockedCategory = "video_tape_rental_stores"
	IssuingCardholderSpendingControlsBlockedCategoryVocationalTradeSchools                                     IssuingCardholderSpendingControlsBlockedCategory = "vocational_trade_schools"
	IssuingCardholderSpendingControlsBlockedCategoryWatchJewelryRepair                                         IssuingCardholderSpendingControlsBlockedCategory = "watch_jewelry_repair"
	IssuingCardholderSpendingControlsBlockedCategoryWeldingRepair                                              IssuingCardholderSpendingControlsBlockedCategory = "welding_repair"
	IssuingCardholderSpendingControlsBlockedCategoryWholesaleClubs                                             IssuingCardholderSpendingControlsBlockedCategory = "wholesale_clubs"
	IssuingCardholderSpendingControlsBlockedCategoryWigAndToupeeStores                                         IssuingCardholderSpendingControlsBlockedCategory = "wig_and_toupee_stores"
	IssuingCardholderSpendingControlsBlockedCategoryWiresMoneyOrders                                           IssuingCardholderSpendingControlsBlockedCategory = "wires_money_orders"
	IssuingCardholderSpendingControlsBlockedCategoryWomensAccessoryAndSpecialtyShops                           IssuingCardholderSpendingControlsBlockedCategory = "womens_accessory_and_specialty_shops"
	IssuingCardholderSpendingControlsBlockedCategoryWomensReadyToWearStores                                    IssuingCardholderSpendingControlsBlockedCategory = "womens_ready_to_wear_stores"
	IssuingCardholderSpendingControlsBlockedCategoryWreckingAndSalvageYards                                    IssuingCardholderSpendingControlsBlockedCategory = "wrecking_and_salvage_yards"
)

type IssuingCardholderSpendingControlsSpendingLimitCategory string

const (
	IssuingCardholderSpendingControlsSpendingLimitCategoryAcRefrigerationRepair                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "ac_refrigeration_repair"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAccountingBookkeepingServices                              IssuingCardholderSpendingControlsSpendingLimitCategory = "accounting_bookkeeping_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAdvertisingServices                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "advertising_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAgriculturalCooperative                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "agricultural_cooperative"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAirlinesAirCarriers                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "airlines_air_carriers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAirportsFlyingFields                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "airports_flying_fields"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAmbulanceServices                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "ambulance_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAmusementParksCarnivals                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "amusement_parks_carnivals"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAntiqueReproductions                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "antique_reproductions"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAntiqueShops                                               IssuingCardholderSpendingControlsSpendingLimitCategory = "antique_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAquariums                                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "aquariums"
	IssuingCardholderSpendingControlsSpendingLimitCategoryArchitecturalSurveyingServices                             IssuingCardholderSpendingControlsSpendingLimitCategory = "architectural_surveying_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryArtDealersAndGalleries                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "art_dealers_and_galleries"
	IssuingCardholderSpendingControlsSpendingLimitCategoryArtistsSupplyAndCraftShops                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "artists_supply_and_craft_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAutoAndHomeSupplyStores                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "auto_and_home_supply_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAutoBodyRepairShops                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "auto_body_repair_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAutoPaintShops                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "auto_paint_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAutoServiceShops                                           IssuingCardholderSpendingControlsSpendingLimitCategory = "auto_service_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAutomatedCashDisburse                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "automated_cash_disburse"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAutomatedFuelDispensers                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "automated_fuel_dispensers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAutomobileAssociations                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "automobile_associations"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAutomotivePartsAndAccessoriesStores                        IssuingCardholderSpendingControlsSpendingLimitCategory = "automotive_parts_and_accessories_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryAutomotiveTireStores                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "automotive_tire_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBailAndBondPayments                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "bail_and_bond_payments"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBakeries                                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "bakeries"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBandsOrchestras                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "bands_orchestras"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBarberAndBeautyShops                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "barber_and_beauty_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBettingCasinoGambling                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "betting_casino_gambling"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBicycleShops                                               IssuingCardholderSpendingControlsSpendingLimitCategory = "bicycle_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBilliardPoolEstablishments                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "billiard_pool_establishments"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBoatDealers                                                IssuingCardholderSpendingControlsSpendingLimitCategory = "boat_dealers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBoatRentalsAndLeases                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "boat_rentals_and_leases"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBookStores                                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "book_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBooksPeriodicalsAndNewspapers                              IssuingCardholderSpendingControlsSpendingLimitCategory = "books_periodicals_and_newspapers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBowlingAlleys                                              IssuingCardholderSpendingControlsSpendingLimitCategory = "bowling_alleys"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBusLines                                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "bus_lines"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBusinessSecretarialSchools                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "business_secretarial_schools"
	IssuingCardholderSpendingControlsSpendingLimitCategoryBuyingShoppingServices                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "buying_shopping_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCableSatelliteAndOtherPayTelevisionAndRadio                IssuingCardholderSpendingControlsSpendingLimitCategory = "cable_satellite_and_other_pay_television_and_radio"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCameraAndPhotographicSupplyStores                          IssuingCardholderSpendingControlsSpendingLimitCategory = "camera_and_photographic_supply_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCandyNutAndConfectioneryStores                             IssuingCardholderSpendingControlsSpendingLimitCategory = "candy_nut_and_confectionery_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCarAndTruckDealersNewUsed                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "car_and_truck_dealers_new_used"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCarAndTruckDealersUsedOnly                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "car_and_truck_dealers_used_only"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCarRentalAgencies                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "car_rental_agencies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCarWashes                                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "car_washes"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCarpentryServices                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "carpentry_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCarpetUpholsteryCleaning                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "carpet_upholstery_cleaning"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCaterers                                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "caterers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCharitableAndSocialServiceOrganizationsFundraising         IssuingCardholderSpendingControlsSpendingLimitCategory = "charitable_and_social_service_organizations_fundraising"
	IssuingCardholderSpendingControlsSpendingLimitCategoryChemicalsAndAlliedProducts                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "chemicals_and_allied_products"
	IssuingCardholderSpendingControlsSpendingLimitCategoryChildCareServices                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "child_care_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryChildrensAndInfantsWearStores                              IssuingCardholderSpendingControlsSpendingLimitCategory = "childrens_and_infants_wear_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryChiropodistsPodiatrists                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "chiropodists_podiatrists"
	IssuingCardholderSpendingControlsSpendingLimitCategoryChiropractors                                              IssuingCardholderSpendingControlsSpendingLimitCategory = "chiropractors"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCigarStoresAndStands                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "cigar_stores_and_stands"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCivicSocialFraternalAssociations                           IssuingCardholderSpendingControlsSpendingLimitCategory = "civic_social_fraternal_associations"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCleaningAndMaintenance                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "cleaning_and_maintenance"
	IssuingCardholderSpendingControlsSpendingLimitCategoryClothingRental                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "clothing_rental"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCollegesUniversities                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "colleges_universities"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCommercialEquipment                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "commercial_equipment"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCommercialFootwear                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "commercial_footwear"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCommercialPhotographyArtAndGraphics                        IssuingCardholderSpendingControlsSpendingLimitCategory = "commercial_photography_art_and_graphics"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCommuterTransportAndFerries                                IssuingCardholderSpendingControlsSpendingLimitCategory = "commuter_transport_and_ferries"
	IssuingCardholderSpendingControlsSpendingLimitCategoryComputerNetworkServices                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "computer_network_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryComputerProgramming                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "computer_programming"
	IssuingCardholderSpendingControlsSpendingLimitCategoryComputerRepair                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "computer_repair"
	IssuingCardholderSpendingControlsSpendingLimitCategoryComputerSoftwareStores                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "computer_software_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryComputersPeripheralsAndSoftware                            IssuingCardholderSpendingControlsSpendingLimitCategory = "computers_peripherals_and_software"
	IssuingCardholderSpendingControlsSpendingLimitCategoryConcreteWorkServices                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "concrete_work_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryConstructionMaterials                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "construction_materials"
	IssuingCardholderSpendingControlsSpendingLimitCategoryConsultingPublicRelations                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "consulting_public_relations"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCorrespondenceSchools                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "correspondence_schools"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCosmeticStores                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "cosmetic_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCounselingServices                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "counseling_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCountryClubs                                               IssuingCardholderSpendingControlsSpendingLimitCategory = "country_clubs"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCourierServices                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "courier_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCourtCosts                                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "court_costs"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCreditReportingAgencies                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "credit_reporting_agencies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryCruiseLines                                                IssuingCardholderSpendingControlsSpendingLimitCategory = "cruise_lines"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDairyProductsStores                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "dairy_products_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDanceHallStudiosSchools                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "dance_hall_studios_schools"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDatingEscortServices                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "dating_escort_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDentistsOrthodontists                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "dentists_orthodontists"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDepartmentStores                                           IssuingCardholderSpendingControlsSpendingLimitCategory = "department_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDetectiveAgencies                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "detective_agencies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDigitalGoodsApplications                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "digital_goods_applications"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDigitalGoodsGames                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "digital_goods_games"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDigitalGoodsLargeVolume                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "digital_goods_large_volume"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDigitalGoodsMedia                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "digital_goods_media"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDirectMarketingCatalogMerchant                             IssuingCardholderSpendingControlsSpendingLimitCategory = "direct_marketing_catalog_merchant"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDirectMarketingCombinationCatalogAndRetailMerchant         IssuingCardholderSpendingControlsSpendingLimitCategory = "direct_marketing_combination_catalog_and_retail_merchant"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDirectMarketingInboundTelemarketing                        IssuingCardholderSpendingControlsSpendingLimitCategory = "direct_marketing_inbound_telemarketing"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDirectMarketingInsuranceServices                           IssuingCardholderSpendingControlsSpendingLimitCategory = "direct_marketing_insurance_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDirectMarketingOther                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "direct_marketing_other"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDirectMarketingOutboundTelemarketing                       IssuingCardholderSpendingControlsSpendingLimitCategory = "direct_marketing_outbound_telemarketing"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDirectMarketingSubscription                                IssuingCardholderSpendingControlsSpendingLimitCategory = "direct_marketing_subscription"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDirectMarketingTravel                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "direct_marketing_travel"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDiscountStores                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "discount_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDoctors                                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "doctors"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDoorToDoorSales                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "door_to_door_sales"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDraperyWindowCoveringAndUpholsteryStores                   IssuingCardholderSpendingControlsSpendingLimitCategory = "drapery_window_covering_and_upholstery_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDrinkingPlaces                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "drinking_places"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDrugStoresAndPharmacies                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "drug_stores_and_pharmacies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDrugsDrugProprietariesAndDruggistSundries                  IssuingCardholderSpendingControlsSpendingLimitCategory = "drugs_drug_proprietaries_and_druggist_sundries"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDryCleaners                                                IssuingCardholderSpendingControlsSpendingLimitCategory = "dry_cleaners"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDurableGoods                                               IssuingCardholderSpendingControlsSpendingLimitCategory = "durable_goods"
	IssuingCardholderSpendingControlsSpendingLimitCategoryDutyFreeStores                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "duty_free_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryEatingPlacesRestaurants                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "eating_places_restaurants"
	IssuingCardholderSpendingControlsSpendingLimitCategoryEducationalServices                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "educational_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryElectricRazorStores                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "electric_razor_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryElectricalPartsAndEquipment                                IssuingCardholderSpendingControlsSpendingLimitCategory = "electrical_parts_and_equipment"
	IssuingCardholderSpendingControlsSpendingLimitCategoryElectricalServices                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "electrical_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryElectronicsRepairShops                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "electronics_repair_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryElectronicsStores                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "electronics_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryElementarySecondarySchools                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "elementary_secondary_schools"
	IssuingCardholderSpendingControlsSpendingLimitCategoryEmploymentTempAgencies                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "employment_temp_agencies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryEquipmentRental                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "equipment_rental"
	IssuingCardholderSpendingControlsSpendingLimitCategoryExterminatingServices                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "exterminating_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFamilyClothingStores                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "family_clothing_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFastFoodRestaurants                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "fast_food_restaurants"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFinancialInstitutions                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "financial_institutions"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFinesGovernmentAdministrativeEntities                      IssuingCardholderSpendingControlsSpendingLimitCategory = "fines_government_administrative_entities"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFireplaceFireplaceScreensAndAccessoriesStores              IssuingCardholderSpendingControlsSpendingLimitCategory = "fireplace_fireplace_screens_and_accessories_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFloorCoveringStores                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "floor_covering_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFlorists                                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "florists"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFloristsSuppliesNurseryStockAndFlowers                     IssuingCardholderSpendingControlsSpendingLimitCategory = "florists_supplies_nursery_stock_and_flowers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFreezerAndLockerMeatProvisioners                           IssuingCardholderSpendingControlsSpendingLimitCategory = "freezer_and_locker_meat_provisioners"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFuelDealersNonAutomotive                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "fuel_dealers_non_automotive"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFuneralServicesCrematories                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "funeral_services_crematories"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFurnitureHomeFurnishingsAndEquipmentStoresExceptAppliances IssuingCardholderSpendingControlsSpendingLimitCategory = "furniture_home_furnishings_and_equipment_stores_except_appliances"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFurnitureRepairRefinishing                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "furniture_repair_refinishing"
	IssuingCardholderSpendingControlsSpendingLimitCategoryFurriersAndFurShops                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "furriers_and_fur_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryGeneralServices                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "general_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryGiftCardNoveltyAndSouvenirShops                            IssuingCardholderSpendingControlsSpendingLimitCategory = "gift_card_novelty_and_souvenir_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryGlassPaintAndWallpaperStores                               IssuingCardholderSpendingControlsSpendingLimitCategory = "glass_paint_and_wallpaper_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryGlasswareCrystalStores                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "glassware_crystal_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryGolfCoursesPublic                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "golf_courses_public"
	IssuingCardholderSpendingControlsSpendingLimitCategoryGovernmentServices                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "government_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryGroceryStoresSupermarkets                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "grocery_stores_supermarkets"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHardwareEquipmentAndSupplies                               IssuingCardholderSpendingControlsSpendingLimitCategory = "hardware_equipment_and_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHardwareStores                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "hardware_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHealthAndBeautySpas                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "health_and_beauty_spas"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHearingAidsSalesAndSupplies                                IssuingCardholderSpendingControlsSpendingLimitCategory = "hearing_aids_sales_and_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHeatingPlumbingAC                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "heating_plumbing_a_c"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHobbyToyAndGameShops                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "hobby_toy_and_game_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHomeSupplyWarehouseStores                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "home_supply_warehouse_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHospitals                                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "hospitals"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHotelsMotelsAndResorts                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "hotels_motels_and_resorts"
	IssuingCardholderSpendingControlsSpendingLimitCategoryHouseholdApplianceStores                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "household_appliance_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryIndustrialSupplies                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "industrial_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryInformationRetrievalServices                               IssuingCardholderSpendingControlsSpendingLimitCategory = "information_retrieval_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryInsuranceDefault                                           IssuingCardholderSpendingControlsSpendingLimitCategory = "insurance_default"
	IssuingCardholderSpendingControlsSpendingLimitCategoryInsuranceUnderwritingPremiums                              IssuingCardholderSpendingControlsSpendingLimitCategory = "insurance_underwriting_premiums"
	IssuingCardholderSpendingControlsSpendingLimitCategoryIntraCompanyPurchases                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "intra_company_purchases"
	IssuingCardholderSpendingControlsSpendingLimitCategoryJewelryStoresWatchesClocksAndSilverwareStores              IssuingCardholderSpendingControlsSpendingLimitCategory = "jewelry_stores_watches_clocks_and_silverware_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryLandscapingServices                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "landscaping_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryLaundries                                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "laundries"
	IssuingCardholderSpendingControlsSpendingLimitCategoryLaundryCleaningServices                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "laundry_cleaning_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryLegalServicesAttorneys                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "legal_services_attorneys"
	IssuingCardholderSpendingControlsSpendingLimitCategoryLuggageAndLeatherGoodsStores                               IssuingCardholderSpendingControlsSpendingLimitCategory = "luggage_and_leather_goods_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryLumberBuildingMaterialsStores                              IssuingCardholderSpendingControlsSpendingLimitCategory = "lumber_building_materials_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryManualCashDisburse                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "manual_cash_disburse"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMarinasServiceAndSupplies                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "marinas_service_and_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMasonryStoneworkAndPlaster                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "masonry_stonework_and_plaster"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMassageParlors                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "massage_parlors"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMedicalAndDentalLabs                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "medical_and_dental_labs"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMedicalDentalOphthalmicAndHospitalEquipmentAndSupplies     IssuingCardholderSpendingControlsSpendingLimitCategory = "medical_dental_ophthalmic_and_hospital_equipment_and_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMedicalServices                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "medical_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMembershipOrganizations                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "membership_organizations"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMensAndBoysClothingAndAccessoriesStores                    IssuingCardholderSpendingControlsSpendingLimitCategory = "mens_and_boys_clothing_and_accessories_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMensWomensClothingStores                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "mens_womens_clothing_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMetalServiceCenters                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "metal_service_centers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneous                                              IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousApparelAndAccessoryShops                      IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_apparel_and_accessory_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousAutoDealers                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_auto_dealers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousBusinessServices                              IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_business_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousFoodStores                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_food_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousGeneralMerchandise                            IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_general_merchandise"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousGeneralServices                               IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_general_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousHomeFurnishingSpecialtyStores                 IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_home_furnishing_specialty_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousPublishingAndPrinting                         IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_publishing_and_printing"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousRecreationServices                            IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_recreation_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousRepairShops                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_repair_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMiscellaneousSpecialtyRetail                               IssuingCardholderSpendingControlsSpendingLimitCategory = "miscellaneous_specialty_retail"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMobileHomeDealers                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "mobile_home_dealers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMotionPictureTheaters                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "motion_picture_theaters"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMotorFreightCarriersAndTrucking                            IssuingCardholderSpendingControlsSpendingLimitCategory = "motor_freight_carriers_and_trucking"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMotorHomesDealers                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "motor_homes_dealers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMotorVehicleSuppliesAndNewParts                            IssuingCardholderSpendingControlsSpendingLimitCategory = "motor_vehicle_supplies_and_new_parts"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMotorcycleShopsAndDealers                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "motorcycle_shops_and_dealers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMotorcycleShopsDealers                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "motorcycle_shops_dealers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryMusicStoresMusicalInstrumentsPianosAndSheetMusic           IssuingCardholderSpendingControlsSpendingLimitCategory = "music_stores_musical_instruments_pianos_and_sheet_music"
	IssuingCardholderSpendingControlsSpendingLimitCategoryNewsDealersAndNewsstands                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "news_dealers_and_newsstands"
	IssuingCardholderSpendingControlsSpendingLimitCategoryNonFiMoneyOrders                                           IssuingCardholderSpendingControlsSpendingLimitCategory = "non_fi_money_orders"
	IssuingCardholderSpendingControlsSpendingLimitCategoryNonFiStoredValueCardPurchaseLoad                           IssuingCardholderSpendingControlsSpendingLimitCategory = "non_fi_stored_value_card_purchase_load"
	IssuingCardholderSpendingControlsSpendingLimitCategoryNondurableGoods                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "nondurable_goods"
	IssuingCardholderSpendingControlsSpendingLimitCategoryNurseriesLawnAndGardenSupplyStores                         IssuingCardholderSpendingControlsSpendingLimitCategory = "nurseries_lawn_and_garden_supply_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryNursingPersonalCare                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "nursing_personal_care"
	IssuingCardholderSpendingControlsSpendingLimitCategoryOfficeAndCommercialFurniture                               IssuingCardholderSpendingControlsSpendingLimitCategory = "office_and_commercial_furniture"
	IssuingCardholderSpendingControlsSpendingLimitCategoryOpticiansEyeglasses                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "opticians_eyeglasses"
	IssuingCardholderSpendingControlsSpendingLimitCategoryOptometristsOphthalmologist                                IssuingCardholderSpendingControlsSpendingLimitCategory = "optometrists_ophthalmologist"
	IssuingCardholderSpendingControlsSpendingLimitCategoryOrthopedicGoodsProstheticDevices                           IssuingCardholderSpendingControlsSpendingLimitCategory = "orthopedic_goods_prosthetic_devices"
	IssuingCardholderSpendingControlsSpendingLimitCategoryOsteopaths                                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "osteopaths"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPackageStoresBeerWineAndLiquor                             IssuingCardholderSpendingControlsSpendingLimitCategory = "package_stores_beer_wine_and_liquor"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPaintsVarnishesAndSupplies                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "paints_varnishes_and_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryParkingLotsGarages                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "parking_lots_garages"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPassengerRailways                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "passenger_railways"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPawnShops                                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "pawn_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPetShopsPetFoodAndSupplies                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "pet_shops_pet_food_and_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPetroleumAndPetroleumProducts                              IssuingCardholderSpendingControlsSpendingLimitCategory = "petroleum_and_petroleum_products"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPhotoDeveloping                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "photo_developing"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPhotographicPhotocopyMicrofilmEquipmentAndSupplies         IssuingCardholderSpendingControlsSpendingLimitCategory = "photographic_photocopy_microfilm_equipment_and_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPhotographicStudios                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "photographic_studios"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPictureVideoProduction                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "picture_video_production"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPieceGoodsNotionsAndOtherDryGoods                          IssuingCardholderSpendingControlsSpendingLimitCategory = "piece_goods_notions_and_other_dry_goods"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPlumbingHeatingEquipmentAndSupplies                        IssuingCardholderSpendingControlsSpendingLimitCategory = "plumbing_heating_equipment_and_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPoliticalOrganizations                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "political_organizations"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPostalServicesGovernmentOnly                               IssuingCardholderSpendingControlsSpendingLimitCategory = "postal_services_government_only"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPreciousStonesAndMetalsWatchesAndJewelry                   IssuingCardholderSpendingControlsSpendingLimitCategory = "precious_stones_and_metals_watches_and_jewelry"
	IssuingCardholderSpendingControlsSpendingLimitCategoryProfessionalServices                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "professional_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryPublicWarehousingAndStorage                                IssuingCardholderSpendingControlsSpendingLimitCategory = "public_warehousing_and_storage"
	IssuingCardholderSpendingControlsSpendingLimitCategoryQuickCopyReproAndBlueprint                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "quick_copy_repro_and_blueprint"
	IssuingCardholderSpendingControlsSpendingLimitCategoryRailroads                                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "railroads"
	IssuingCardholderSpendingControlsSpendingLimitCategoryRealEstateAgentsAndManagersRentals                         IssuingCardholderSpendingControlsSpendingLimitCategory = "real_estate_agents_and_managers_rentals"
	IssuingCardholderSpendingControlsSpendingLimitCategoryRecordStores                                               IssuingCardholderSpendingControlsSpendingLimitCategory = "record_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryRecreationalVehicleRentals                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "recreational_vehicle_rentals"
	IssuingCardholderSpendingControlsSpendingLimitCategoryReligiousGoodsStores                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "religious_goods_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryReligiousOrganizations                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "religious_organizations"
	IssuingCardholderSpendingControlsSpendingLimitCategoryRoofingSidingSheetMetal                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "roofing_siding_sheet_metal"
	IssuingCardholderSpendingControlsSpendingLimitCategorySecretarialSupportServices                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "secretarial_support_services"
	IssuingCardholderSpendingControlsSpendingLimitCategorySecurityBrokersDealers                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "security_brokers_dealers"
	IssuingCardholderSpendingControlsSpendingLimitCategoryServiceStations                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "service_stations"
	IssuingCardholderSpendingControlsSpendingLimitCategorySewingNeedleworkFabricAndPieceGoodsStores                  IssuingCardholderSpendingControlsSpendingLimitCategory = "sewing_needlework_fabric_and_piece_goods_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryShoeRepairHatCleaning                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "shoe_repair_hat_cleaning"
	IssuingCardholderSpendingControlsSpendingLimitCategoryShoeStores                                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "shoe_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategorySmallApplianceRepair                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "small_appliance_repair"
	IssuingCardholderSpendingControlsSpendingLimitCategorySnowmobileDealers                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "snowmobile_dealers"
	IssuingCardholderSpendingControlsSpendingLimitCategorySpecialTradeServices                                       IssuingCardholderSpendingControlsSpendingLimitCategory = "special_trade_services"
	IssuingCardholderSpendingControlsSpendingLimitCategorySpecialtyCleaning                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "specialty_cleaning"
	IssuingCardholderSpendingControlsSpendingLimitCategorySportingGoodsStores                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "sporting_goods_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategorySportingRecreationCamps                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "sporting_recreation_camps"
	IssuingCardholderSpendingControlsSpendingLimitCategorySportsAndRidingApparelStores                               IssuingCardholderSpendingControlsSpendingLimitCategory = "sports_and_riding_apparel_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategorySportsClubsFields                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "sports_clubs_fields"
	IssuingCardholderSpendingControlsSpendingLimitCategoryStampAndCoinStores                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "stamp_and_coin_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryStationaryOfficeSuppliesPrintingAndWritingPaper            IssuingCardholderSpendingControlsSpendingLimitCategory = "stationary_office_supplies_printing_and_writing_paper"
	IssuingCardholderSpendingControlsSpendingLimitCategoryStationeryStoresOfficeAndSchoolSupplyStores                IssuingCardholderSpendingControlsSpendingLimitCategory = "stationery_stores_office_and_school_supply_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategorySwimmingPoolsSales                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "swimming_pools_sales"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTUiTravelGermany                                           IssuingCardholderSpendingControlsSpendingLimitCategory = "t_ui_travel_germany"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTailorsAlterations                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "tailors_alterations"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTaxPaymentsGovernmentAgencies                              IssuingCardholderSpendingControlsSpendingLimitCategory = "tax_payments_government_agencies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTaxPreparationServices                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "tax_preparation_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTaxicabsLimousines                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "taxicabs_limousines"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTelecommunicationEquipmentAndTelephoneSales                IssuingCardholderSpendingControlsSpendingLimitCategory = "telecommunication_equipment_and_telephone_sales"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTelecommunicationServices                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "telecommunication_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTelegraphServices                                          IssuingCardholderSpendingControlsSpendingLimitCategory = "telegraph_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTentAndAwningShops                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "tent_and_awning_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTestingLaboratories                                        IssuingCardholderSpendingControlsSpendingLimitCategory = "testing_laboratories"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTheatricalTicketAgencies                                   IssuingCardholderSpendingControlsSpendingLimitCategory = "theatrical_ticket_agencies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTimeshares                                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "timeshares"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTireRetreadingAndRepair                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "tire_retreading_and_repair"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTollsBridgeFees                                            IssuingCardholderSpendingControlsSpendingLimitCategory = "tolls_bridge_fees"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTouristAttractionsAndExhibits                              IssuingCardholderSpendingControlsSpendingLimitCategory = "tourist_attractions_and_exhibits"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTowingServices                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "towing_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTrailerParksCampgrounds                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "trailer_parks_campgrounds"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTransportationServices                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "transportation_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTravelAgenciesTourOperators                                IssuingCardholderSpendingControlsSpendingLimitCategory = "travel_agencies_tour_operators"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTruckStopIteration                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "truck_stop_iteration"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTruckUtilityTrailerRentals                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "truck_utility_trailer_rentals"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTypesettingPlateMakingAndRelatedServices                   IssuingCardholderSpendingControlsSpendingLimitCategory = "typesetting_plate_making_and_related_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryTypewriterStores                                           IssuingCardholderSpendingControlsSpendingLimitCategory = "typewriter_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryUSFederalGovernmentAgenciesOrDepartments                   IssuingCardholderSpendingControlsSpendingLimitCategory = "u_s_federal_government_agencies_or_departments"
	IssuingCardholderSpendingControlsSpendingLimitCategoryUniformsCommercialClothing                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "uniforms_commercial_clothing"
	IssuingCardholderSpendingControlsSpendingLimitCategoryUsedMerchandiseAndSecondhandStores                         IssuingCardholderSpendingControlsSpendingLimitCategory = "used_merchandise_and_secondhand_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryUtilities                                                  IssuingCardholderSpendingControlsSpendingLimitCategory = "utilities"
	IssuingCardholderSpendingControlsSpendingLimitCategoryVarietyStores                                              IssuingCardholderSpendingControlsSpendingLimitCategory = "variety_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryVeterinaryServices                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "veterinary_services"
	IssuingCardholderSpendingControlsSpendingLimitCategoryVideoAmusementGameSupplies                                 IssuingCardholderSpendingControlsSpendingLimitCategory = "video_amusement_game_supplies"
	IssuingCardholderSpendingControlsSpendingLimitCategoryVideoGameArcades                                           IssuingCardholderSpendingControlsSpendingLimitCategory = "video_game_arcades"
	IssuingCardholderSpendingControlsSpendingLimitCategoryVideoTapeRentalStores                                      IssuingCardholderSpendingControlsSpendingLimitCategory = "video_tape_rental_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryVocationalTradeSchools                                     IssuingCardholderSpendingControlsSpendingLimitCategory = "vocational_trade_schools"
	IssuingCardholderSpendingControlsSpendingLimitCategoryWatchJewelryRepair                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "watch_jewelry_repair"
	IssuingCardholderSpendingControlsSpendingLimitCategoryWeldingRepair                                              IssuingCardholderSpendingControlsSpendingLimitCategory = "welding_repair"
	IssuingCardholderSpendingControlsSpendingLimitCategoryWholesaleClubs                                             IssuingCardholderSpendingControlsSpendingLimitCategory = "wholesale_clubs"
	IssuingCardholderSpendingControlsSpendingLimitCategoryWigAndToupeeStores                                         IssuingCardholderSpendingControlsSpendingLimitCategory = "wig_and_toupee_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryWiresMoneyOrders                                           IssuingCardholderSpendingControlsSpendingLimitCategory = "wires_money_orders"
	IssuingCardholderSpendingControlsSpendingLimitCategoryWomensAccessoryAndSpecialtyShops                           IssuingCardholderSpendingControlsSpendingLimitCategory = "womens_accessory_and_specialty_shops"
	IssuingCardholderSpendingControlsSpendingLimitCategoryWomensReadyToWearStores                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "womens_ready_to_wear_stores"
	IssuingCardholderSpendingControlsSpendingLimitCategoryWreckingAndSalvageYards                                    IssuingCardholderSpendingControlsSpendingLimitCategory = "wrecking_and_salvage_yards"
)

// IssuingCardholderSpendingControlsSpendingLimitInterval is the list of possible values for the interval
// for a spending limit on an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimitInterval string

// List of values that IssuingCardShippingStatus can take.
const (
	IssuingCardholderSpendingControlsSpendingLimitIntervalAllTime          IssuingCardholderSpendingControlsSpendingLimitInterval = "all_time"
	IssuingCardholderSpendingControlsSpendingLimitIntervalDaily            IssuingCardholderSpendingControlsSpendingLimitInterval = "daily"
	IssuingCardholderSpendingControlsSpendingLimitIntervalMonthly          IssuingCardholderSpendingControlsSpendingLimitInterval = "monthly"
	IssuingCardholderSpendingControlsSpendingLimitIntervalPerAuthorization IssuingCardholderSpendingControlsSpendingLimitInterval = "per_authorization"
	IssuingCardholderSpendingControlsSpendingLimitIntervalWeekly           IssuingCardholderSpendingControlsSpendingLimitInterval = "weekly"
	IssuingCardholderSpendingControlsSpendingLimitIntervalYearly           IssuingCardholderSpendingControlsSpendingLimitInterval = "yearly"
)

// IssuingCardholderStatus is the possible values for status on an issuing cardholder.
type IssuingCardholderStatus string

// List of values that IssuingCardholderStatus can take.
const (
	IssuingCardholderStatusActive   IssuingCardholderStatus = "active"
	IssuingCardholderStatusBlocked  IssuingCardholderStatus = "blocked"
	IssuingCardholderStatusInactive IssuingCardholderStatus = "inactive"
)

// IssuingCardholderType is the type of an issuing cardholder.
type IssuingCardholderType string

// List of values that IssuingCardholderType can take.
const (
	IssuingCardholderTypeCompany    IssuingCardholderType = "company"
	IssuingCardholderTypeIndividual IssuingCardholderType = "individual"
)

// IssuingCardholderBillingParams is the set of parameters that can be used for billing with the Issuing APIs.
type IssuingCardholderBillingParams struct {
	Address *AddressParams `form:"address"`
}

// IssuingCardholderCompanyParams represents additional information about a company cardholder.
type IssuingCardholderCompanyParams struct {
	TaxID *string `form:"tax_id"`
}
type IssuingCardholderIndividualDobParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// IssuingCardholderIndividualVerificationDocumentParams represents an
// identifying document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// IssuingCardholderIndividualVerificationParams represents government-issued ID
// document for this cardholder.
type IssuingCardholderIndividualVerificationParams struct {
	Document *IssuingCardholderIndividualVerificationDocumentParams `form:"document"`
}

// IssuingCardholderIndividualParams represents additional information about an
// `individual` cardholder.
type IssuingCardholderIndividualParams struct {
	Dob          *IssuingCardholderIndividualDobParams          `form:"dob"`
	FirstName    *string                                        `form:"first_name"`
	LastName     *string                                        `form:"last_name"`
	Verification *IssuingCardholderIndividualVerificationParams `form:"verification"`
}

// IssuingCardholderSpendingControlsSpendingLimitParams is the set of parameters that can be used to
// represent a given spending limit for an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimitParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// IssuingCardholderSpendingControlsParams is the set of parameters that can be used to configure
// the spending controls for an issuing cardholder
type IssuingCardholderSpendingControlsParams struct {
	AllowedCategories      []*string                                               `form:"allowed_categories"`
	BlockedCategories      []*string                                               `form:"blocked_categories"`
	SpendingLimits         []*IssuingCardholderSpendingControlsSpendingLimitParams `form:"spending_limits"`
	SpendingLimitsCurrency *string                                                 `form:"spending_limits_currency"`
}

// IssuingCardholderParams is the set of parameters that can be used when creating or updating an issuing cardholder.
type IssuingCardholderParams struct {
	Params           `form:"*"`
	Billing          *IssuingCardholderBillingParams          `form:"billing"`
	Company          *IssuingCardholderCompanyParams          `form:"company"`
	Email            *string                                  `form:"email"`
	Individual       *IssuingCardholderIndividualParams       `form:"individual"`
	Name             *string                                  `form:"name"`
	PhoneNumber      *string                                  `form:"phone_number"`
	SpendingControls *IssuingCardholderSpendingControlsParams `form:"spending_controls"`
	Status           *string                                  `form:"status"`
	Type             *string                                  `form:"type"`
}

// IssuingCardholderListParams is the set of parameters that can be used when listing issuing cardholders.
type IssuingCardholderListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Email        *string           `form:"email"`
	PhoneNumber  *string           `form:"phone_number"`
	Status       *string           `form:"status"`
	Type         *string           `form:"type"`
}

// IssuingCardholderBilling is the resource representing the billing hash with the Issuing APIs.
type IssuingCardholderBilling struct {
	Address Address `json:"address"`
}

// IssuingCardholderCompany represents additional information about a company cardholder.
type IssuingCardholderCompany struct {
	TaxIDProvided bool `json:"tax_id_provided"`
}
type IssuingCardholderIndividualDob struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// IssuingCardholderIndividualVerificationDocument represents an identifying
// document, either a passport or local ID card.
type IssuingCardholderIndividualVerificationDocument struct {
	Back  *File `json:"back"`
	Front *File `json:"front"`
}

// IssuingCardholderIndividualVerification represents the Government-issued ID
// document for this cardholder
type IssuingCardholderIndividualVerification struct {
	Document *IssuingCardholderIndividualVerificationDocument `json:"document"`
}

// IssuingCardholderIndividual represents additional information about an
// individual cardholder.
type IssuingCardholderIndividual struct {
	Dob          *IssuingCardholderIndividualDob          `json:"dob"`
	FirstName    string                                   `json:"first_name"`
	LastName     string                                   `json:"last_name"`
	Verification *IssuingCardholderIndividualVerification `json:"verification"`
}

// IssuingCardholderRequirements contains the verification requirements for the cardholder.
type IssuingCardholderRequirements struct {
	DisabledReason IssuingCardholderRequirementsDisabledReason `json:"disabled_reason"`
	PastDue        []IssuingCardholderRequirementsPastDue      `json:"past_due"`
}

// IssuingCardholderSpendingControlsSpendingLimit is the resource representing a spending limit
// for an issuing cardholder.
type IssuingCardholderSpendingControlsSpendingLimit struct {
	Amount     int64                                                    `json:"amount"`
	Categories []IssuingCardholderSpendingControlsSpendingLimitCategory `json:"categories"`
	Interval   IssuingCardholderSpendingControlsSpendingLimitInterval   `json:"interval"`
}

// IssuingCardholderSpendingControls is the resource representing spending controls
// for an issuing cardholder.
type IssuingCardholderSpendingControls struct {
	AllowedCategories      []IssuingCardholderSpendingControlsAllowedCategory `json:"allowed_categories"`
	BlockedCategories      []IssuingCardholderSpendingControlsBlockedCategory `json:"blocked_categories"`
	SpendingLimits         []*IssuingCardholderSpendingControlsSpendingLimit  `json:"spending_limits"`
	SpendingLimitsCurrency Currency                                           `json:"spending_limits_currency"`
}

// IssuingCardholder is the resource representing a Stripe issuing cardholder.
type IssuingCardholder struct {
	APIResource
	Billing          *IssuingCardholderBilling          `json:"billing"`
	Company          *IssuingCardholderCompany          `json:"company"`
	Created          int64                              `json:"created"`
	Email            string                             `json:"email"`
	ID               string                             `json:"id"`
	Individual       *IssuingCardholderIndividual       `json:"individual"`
	Livemode         bool                               `json:"livemode"`
	Metadata         map[string]string                  `json:"metadata"`
	Name             string                             `json:"name"`
	Object           string                             `json:"object"`
	PhoneNumber      string                             `json:"phone_number"`
	Requirements     *IssuingCardholderRequirements     `json:"requirements"`
	SpendingControls *IssuingCardholderSpendingControls `json:"spending_controls"`
	Status           IssuingCardholderStatus            `json:"status"`
	Type             IssuingCardholderType              `json:"type"`
}

// IssuingCardholderList is a list of issuing cardholders as retrieved from a list endpoint.
type IssuingCardholderList struct {
	APIResource
	ListMeta
	Data []*IssuingCardholder `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingCardholder.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingCardholder) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingCardholder IssuingCardholder
	var v issuingCardholder
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingCardholder(v)
	return nil
}
