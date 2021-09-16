//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// IssuingCardCancellationReason is the list of possible values for the cancellation reason
// on an issuing card.
type IssuingCardCancellationReason string

// List of values that IssuingCardReplacementReason can take.
const (
	IssuingCardCancellationReasonLost   IssuingCardCancellationReason = "lost"
	IssuingCardCancellationReasonStolen IssuingCardCancellationReason = "stolen"
)

// IssuingCardReplacementReason is the list of possible values for the replacement reason on an
// issuing card.
type IssuingCardReplacementReason string

// List of values that IssuingCardReplacementReason can take.
const (
	IssuingCardReplacementReasonDamaged IssuingCardReplacementReason = "damaged"
	IssuingCardReplacementReasonExpired IssuingCardReplacementReason = "expired"
	IssuingCardReplacementReasonLost    IssuingCardReplacementReason = "lost"
	IssuingCardReplacementReasonStolen  IssuingCardReplacementReason = "stolen"
)

// IssuingCardShippingCarrier is the list of possible values for the shipping carrier
// on an issuing card.
type IssuingCardShippingCarrier string

// List of values that IssuingCardShippingCarrier can take.
const (
	IssuingCardShippingCarrierDhl       IssuingCardShippingCarrier = "dhl"
	IssuingCardShippingCarrierFedex     IssuingCardShippingCarrier = "fedex"
	IssuingCardShippingCarrierRoyalMail IssuingCardShippingCarrier = "royal_mail"
	IssuingCardShippingCarrierUsps      IssuingCardShippingCarrier = "usps"
)

// IssuingCardShippingService is the shipment service for a card.
type IssuingCardShippingService string

// List of values that IssuingCardShippingService can take.
const (
	IssuingCardShippingServiceExpress  IssuingCardShippingService = "express"
	IssuingCardShippingServicePriority IssuingCardShippingService = "priority"
	IssuingCardShippingServiceStandard IssuingCardShippingService = "standard"
)

// IssuingCardShippingStatus is the list of possible values for the shipping status
// on an issuing card.
type IssuingCardShippingStatus string

// List of values that IssuingCardShippingStatus can take.
const (
	IssuingCardShippingStatusCanceled  IssuingCardShippingStatus = "canceled"
	IssuingCardShippingStatusDelivered IssuingCardShippingStatus = "delivered"
	IssuingCardShippingStatusFailure   IssuingCardShippingStatus = "failure"
	IssuingCardShippingStatusPending   IssuingCardShippingStatus = "pending"
	IssuingCardShippingStatusReturned  IssuingCardShippingStatus = "returned"
	IssuingCardShippingStatusShipped   IssuingCardShippingStatus = "shipped"
)

// IssuingCardShippingType is the list of possible values for the shipping type
// on an issuing card.
type IssuingCardShippingType string

// List of values that IssuingCardShippingType can take.
const (
	IssuingCardShippingTypeBulk       IssuingCardShippingType = "bulk"
	IssuingCardShippingTypeIndividual IssuingCardShippingType = "individual"
)

// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to allow. All other categories will be blocked. Cannot be set with `blocked_categories`.
type IssuingCardSpendingControlsAllowedCategory string

// List of values that IssuingCardSpendingControlsAllowedCategory can take
const (
	IssuingCardSpendingControlsAllowedCategoryAcRefrigerationRepair                                      IssuingCardSpendingControlsAllowedCategory = "ac_refrigeration_repair"
	IssuingCardSpendingControlsAllowedCategoryAccountingBookkeepingServices                              IssuingCardSpendingControlsAllowedCategory = "accounting_bookkeeping_services"
	IssuingCardSpendingControlsAllowedCategoryAdvertisingServices                                        IssuingCardSpendingControlsAllowedCategory = "advertising_services"
	IssuingCardSpendingControlsAllowedCategoryAgriculturalCooperative                                    IssuingCardSpendingControlsAllowedCategory = "agricultural_cooperative"
	IssuingCardSpendingControlsAllowedCategoryAirlinesAirCarriers                                        IssuingCardSpendingControlsAllowedCategory = "airlines_air_carriers"
	IssuingCardSpendingControlsAllowedCategoryAirportsFlyingFields                                       IssuingCardSpendingControlsAllowedCategory = "airports_flying_fields"
	IssuingCardSpendingControlsAllowedCategoryAmbulanceServices                                          IssuingCardSpendingControlsAllowedCategory = "ambulance_services"
	IssuingCardSpendingControlsAllowedCategoryAmusementParksCarnivals                                    IssuingCardSpendingControlsAllowedCategory = "amusement_parks_carnivals"
	IssuingCardSpendingControlsAllowedCategoryAntiqueReproductions                                       IssuingCardSpendingControlsAllowedCategory = "antique_reproductions"
	IssuingCardSpendingControlsAllowedCategoryAntiqueShops                                               IssuingCardSpendingControlsAllowedCategory = "antique_shops"
	IssuingCardSpendingControlsAllowedCategoryAquariums                                                  IssuingCardSpendingControlsAllowedCategory = "aquariums"
	IssuingCardSpendingControlsAllowedCategoryArchitecturalSurveyingServices                             IssuingCardSpendingControlsAllowedCategory = "architectural_surveying_services"
	IssuingCardSpendingControlsAllowedCategoryArtDealersAndGalleries                                     IssuingCardSpendingControlsAllowedCategory = "art_dealers_and_galleries"
	IssuingCardSpendingControlsAllowedCategoryArtistsSupplyAndCraftShops                                 IssuingCardSpendingControlsAllowedCategory = "artists_supply_and_craft_shops"
	IssuingCardSpendingControlsAllowedCategoryAutoAndHomeSupplyStores                                    IssuingCardSpendingControlsAllowedCategory = "auto_and_home_supply_stores"
	IssuingCardSpendingControlsAllowedCategoryAutoBodyRepairShops                                        IssuingCardSpendingControlsAllowedCategory = "auto_body_repair_shops"
	IssuingCardSpendingControlsAllowedCategoryAutoPaintShops                                             IssuingCardSpendingControlsAllowedCategory = "auto_paint_shops"
	IssuingCardSpendingControlsAllowedCategoryAutoServiceShops                                           IssuingCardSpendingControlsAllowedCategory = "auto_service_shops"
	IssuingCardSpendingControlsAllowedCategoryAutomatedCashDisburse                                      IssuingCardSpendingControlsAllowedCategory = "automated_cash_disburse"
	IssuingCardSpendingControlsAllowedCategoryAutomatedFuelDispensers                                    IssuingCardSpendingControlsAllowedCategory = "automated_fuel_dispensers"
	IssuingCardSpendingControlsAllowedCategoryAutomobileAssociations                                     IssuingCardSpendingControlsAllowedCategory = "automobile_associations"
	IssuingCardSpendingControlsAllowedCategoryAutomotivePartsAndAccessoriesStores                        IssuingCardSpendingControlsAllowedCategory = "automotive_parts_and_accessories_stores"
	IssuingCardSpendingControlsAllowedCategoryAutomotiveTireStores                                       IssuingCardSpendingControlsAllowedCategory = "automotive_tire_stores"
	IssuingCardSpendingControlsAllowedCategoryBailAndBondPayments                                        IssuingCardSpendingControlsAllowedCategory = "bail_and_bond_payments"
	IssuingCardSpendingControlsAllowedCategoryBakeries                                                   IssuingCardSpendingControlsAllowedCategory = "bakeries"
	IssuingCardSpendingControlsAllowedCategoryBandsOrchestras                                            IssuingCardSpendingControlsAllowedCategory = "bands_orchestras"
	IssuingCardSpendingControlsAllowedCategoryBarberAndBeautyShops                                       IssuingCardSpendingControlsAllowedCategory = "barber_and_beauty_shops"
	IssuingCardSpendingControlsAllowedCategoryBettingCasinoGambling                                      IssuingCardSpendingControlsAllowedCategory = "betting_casino_gambling"
	IssuingCardSpendingControlsAllowedCategoryBicycleShops                                               IssuingCardSpendingControlsAllowedCategory = "bicycle_shops"
	IssuingCardSpendingControlsAllowedCategoryBilliardPoolEstablishments                                 IssuingCardSpendingControlsAllowedCategory = "billiard_pool_establishments"
	IssuingCardSpendingControlsAllowedCategoryBoatDealers                                                IssuingCardSpendingControlsAllowedCategory = "boat_dealers"
	IssuingCardSpendingControlsAllowedCategoryBoatRentalsAndLeases                                       IssuingCardSpendingControlsAllowedCategory = "boat_rentals_and_leases"
	IssuingCardSpendingControlsAllowedCategoryBookStores                                                 IssuingCardSpendingControlsAllowedCategory = "book_stores"
	IssuingCardSpendingControlsAllowedCategoryBooksPeriodicalsAndNewspapers                              IssuingCardSpendingControlsAllowedCategory = "books_periodicals_and_newspapers"
	IssuingCardSpendingControlsAllowedCategoryBowlingAlleys                                              IssuingCardSpendingControlsAllowedCategory = "bowling_alleys"
	IssuingCardSpendingControlsAllowedCategoryBusLines                                                   IssuingCardSpendingControlsAllowedCategory = "bus_lines"
	IssuingCardSpendingControlsAllowedCategoryBusinessSecretarialSchools                                 IssuingCardSpendingControlsAllowedCategory = "business_secretarial_schools"
	IssuingCardSpendingControlsAllowedCategoryBuyingShoppingServices                                     IssuingCardSpendingControlsAllowedCategory = "buying_shopping_services"
	IssuingCardSpendingControlsAllowedCategoryCableSatelliteAndOtherPayTelevisionAndRadio                IssuingCardSpendingControlsAllowedCategory = "cable_satellite_and_other_pay_television_and_radio"
	IssuingCardSpendingControlsAllowedCategoryCameraAndPhotographicSupplyStores                          IssuingCardSpendingControlsAllowedCategory = "camera_and_photographic_supply_stores"
	IssuingCardSpendingControlsAllowedCategoryCandyNutAndConfectioneryStores                             IssuingCardSpendingControlsAllowedCategory = "candy_nut_and_confectionery_stores"
	IssuingCardSpendingControlsAllowedCategoryCarAndTruckDealersNewUsed                                  IssuingCardSpendingControlsAllowedCategory = "car_and_truck_dealers_new_used"
	IssuingCardSpendingControlsAllowedCategoryCarAndTruckDealersUsedOnly                                 IssuingCardSpendingControlsAllowedCategory = "car_and_truck_dealers_used_only"
	IssuingCardSpendingControlsAllowedCategoryCarRentalAgencies                                          IssuingCardSpendingControlsAllowedCategory = "car_rental_agencies"
	IssuingCardSpendingControlsAllowedCategoryCarWashes                                                  IssuingCardSpendingControlsAllowedCategory = "car_washes"
	IssuingCardSpendingControlsAllowedCategoryCarpentryServices                                          IssuingCardSpendingControlsAllowedCategory = "carpentry_services"
	IssuingCardSpendingControlsAllowedCategoryCarpetUpholsteryCleaning                                   IssuingCardSpendingControlsAllowedCategory = "carpet_upholstery_cleaning"
	IssuingCardSpendingControlsAllowedCategoryCaterers                                                   IssuingCardSpendingControlsAllowedCategory = "caterers"
	IssuingCardSpendingControlsAllowedCategoryCharitableAndSocialServiceOrganizationsFundraising         IssuingCardSpendingControlsAllowedCategory = "charitable_and_social_service_organizations_fundraising"
	IssuingCardSpendingControlsAllowedCategoryChemicalsAndAlliedProducts                                 IssuingCardSpendingControlsAllowedCategory = "chemicals_and_allied_products"
	IssuingCardSpendingControlsAllowedCategoryChildCareServices                                          IssuingCardSpendingControlsAllowedCategory = "child_care_services"
	IssuingCardSpendingControlsAllowedCategoryChildrensAndInfantsWearStores                              IssuingCardSpendingControlsAllowedCategory = "childrens_and_infants_wear_stores"
	IssuingCardSpendingControlsAllowedCategoryChiropodistsPodiatrists                                    IssuingCardSpendingControlsAllowedCategory = "chiropodists_podiatrists"
	IssuingCardSpendingControlsAllowedCategoryChiropractors                                              IssuingCardSpendingControlsAllowedCategory = "chiropractors"
	IssuingCardSpendingControlsAllowedCategoryCigarStoresAndStands                                       IssuingCardSpendingControlsAllowedCategory = "cigar_stores_and_stands"
	IssuingCardSpendingControlsAllowedCategoryCivicSocialFraternalAssociations                           IssuingCardSpendingControlsAllowedCategory = "civic_social_fraternal_associations"
	IssuingCardSpendingControlsAllowedCategoryCleaningAndMaintenance                                     IssuingCardSpendingControlsAllowedCategory = "cleaning_and_maintenance"
	IssuingCardSpendingControlsAllowedCategoryClothingRental                                             IssuingCardSpendingControlsAllowedCategory = "clothing_rental"
	IssuingCardSpendingControlsAllowedCategoryCollegesUniversities                                       IssuingCardSpendingControlsAllowedCategory = "colleges_universities"
	IssuingCardSpendingControlsAllowedCategoryCommercialEquipment                                        IssuingCardSpendingControlsAllowedCategory = "commercial_equipment"
	IssuingCardSpendingControlsAllowedCategoryCommercialFootwear                                         IssuingCardSpendingControlsAllowedCategory = "commercial_footwear"
	IssuingCardSpendingControlsAllowedCategoryCommercialPhotographyArtAndGraphics                        IssuingCardSpendingControlsAllowedCategory = "commercial_photography_art_and_graphics"
	IssuingCardSpendingControlsAllowedCategoryCommuterTransportAndFerries                                IssuingCardSpendingControlsAllowedCategory = "commuter_transport_and_ferries"
	IssuingCardSpendingControlsAllowedCategoryComputerNetworkServices                                    IssuingCardSpendingControlsAllowedCategory = "computer_network_services"
	IssuingCardSpendingControlsAllowedCategoryComputerProgramming                                        IssuingCardSpendingControlsAllowedCategory = "computer_programming"
	IssuingCardSpendingControlsAllowedCategoryComputerRepair                                             IssuingCardSpendingControlsAllowedCategory = "computer_repair"
	IssuingCardSpendingControlsAllowedCategoryComputerSoftwareStores                                     IssuingCardSpendingControlsAllowedCategory = "computer_software_stores"
	IssuingCardSpendingControlsAllowedCategoryComputersPeripheralsAndSoftware                            IssuingCardSpendingControlsAllowedCategory = "computers_peripherals_and_software"
	IssuingCardSpendingControlsAllowedCategoryConcreteWorkServices                                       IssuingCardSpendingControlsAllowedCategory = "concrete_work_services"
	IssuingCardSpendingControlsAllowedCategoryConstructionMaterials                                      IssuingCardSpendingControlsAllowedCategory = "construction_materials"
	IssuingCardSpendingControlsAllowedCategoryConsultingPublicRelations                                  IssuingCardSpendingControlsAllowedCategory = "consulting_public_relations"
	IssuingCardSpendingControlsAllowedCategoryCorrespondenceSchools                                      IssuingCardSpendingControlsAllowedCategory = "correspondence_schools"
	IssuingCardSpendingControlsAllowedCategoryCosmeticStores                                             IssuingCardSpendingControlsAllowedCategory = "cosmetic_stores"
	IssuingCardSpendingControlsAllowedCategoryCounselingServices                                         IssuingCardSpendingControlsAllowedCategory = "counseling_services"
	IssuingCardSpendingControlsAllowedCategoryCountryClubs                                               IssuingCardSpendingControlsAllowedCategory = "country_clubs"
	IssuingCardSpendingControlsAllowedCategoryCourierServices                                            IssuingCardSpendingControlsAllowedCategory = "courier_services"
	IssuingCardSpendingControlsAllowedCategoryCourtCosts                                                 IssuingCardSpendingControlsAllowedCategory = "court_costs"
	IssuingCardSpendingControlsAllowedCategoryCreditReportingAgencies                                    IssuingCardSpendingControlsAllowedCategory = "credit_reporting_agencies"
	IssuingCardSpendingControlsAllowedCategoryCruiseLines                                                IssuingCardSpendingControlsAllowedCategory = "cruise_lines"
	IssuingCardSpendingControlsAllowedCategoryDairyProductsStores                                        IssuingCardSpendingControlsAllowedCategory = "dairy_products_stores"
	IssuingCardSpendingControlsAllowedCategoryDanceHallStudiosSchools                                    IssuingCardSpendingControlsAllowedCategory = "dance_hall_studios_schools"
	IssuingCardSpendingControlsAllowedCategoryDatingEscortServices                                       IssuingCardSpendingControlsAllowedCategory = "dating_escort_services"
	IssuingCardSpendingControlsAllowedCategoryDentistsOrthodontists                                      IssuingCardSpendingControlsAllowedCategory = "dentists_orthodontists"
	IssuingCardSpendingControlsAllowedCategoryDepartmentStores                                           IssuingCardSpendingControlsAllowedCategory = "department_stores"
	IssuingCardSpendingControlsAllowedCategoryDetectiveAgencies                                          IssuingCardSpendingControlsAllowedCategory = "detective_agencies"
	IssuingCardSpendingControlsAllowedCategoryDigitalGoodsApplications                                   IssuingCardSpendingControlsAllowedCategory = "digital_goods_applications"
	IssuingCardSpendingControlsAllowedCategoryDigitalGoodsGames                                          IssuingCardSpendingControlsAllowedCategory = "digital_goods_games"
	IssuingCardSpendingControlsAllowedCategoryDigitalGoodsLargeVolume                                    IssuingCardSpendingControlsAllowedCategory = "digital_goods_large_volume"
	IssuingCardSpendingControlsAllowedCategoryDigitalGoodsMedia                                          IssuingCardSpendingControlsAllowedCategory = "digital_goods_media"
	IssuingCardSpendingControlsAllowedCategoryDirectMarketingCatalogMerchant                             IssuingCardSpendingControlsAllowedCategory = "direct_marketing_catalog_merchant"
	IssuingCardSpendingControlsAllowedCategoryDirectMarketingCombinationCatalogAndRetailMerchant         IssuingCardSpendingControlsAllowedCategory = "direct_marketing_combination_catalog_and_retail_merchant"
	IssuingCardSpendingControlsAllowedCategoryDirectMarketingInboundTelemarketing                        IssuingCardSpendingControlsAllowedCategory = "direct_marketing_inbound_telemarketing"
	IssuingCardSpendingControlsAllowedCategoryDirectMarketingInsuranceServices                           IssuingCardSpendingControlsAllowedCategory = "direct_marketing_insurance_services"
	IssuingCardSpendingControlsAllowedCategoryDirectMarketingOther                                       IssuingCardSpendingControlsAllowedCategory = "direct_marketing_other"
	IssuingCardSpendingControlsAllowedCategoryDirectMarketingOutboundTelemarketing                       IssuingCardSpendingControlsAllowedCategory = "direct_marketing_outbound_telemarketing"
	IssuingCardSpendingControlsAllowedCategoryDirectMarketingSubscription                                IssuingCardSpendingControlsAllowedCategory = "direct_marketing_subscription"
	IssuingCardSpendingControlsAllowedCategoryDirectMarketingTravel                                      IssuingCardSpendingControlsAllowedCategory = "direct_marketing_travel"
	IssuingCardSpendingControlsAllowedCategoryDiscountStores                                             IssuingCardSpendingControlsAllowedCategory = "discount_stores"
	IssuingCardSpendingControlsAllowedCategoryDoctors                                                    IssuingCardSpendingControlsAllowedCategory = "doctors"
	IssuingCardSpendingControlsAllowedCategoryDoorToDoorSales                                            IssuingCardSpendingControlsAllowedCategory = "door_to_door_sales"
	IssuingCardSpendingControlsAllowedCategoryDraperyWindowCoveringAndUpholsteryStores                   IssuingCardSpendingControlsAllowedCategory = "drapery_window_covering_and_upholstery_stores"
	IssuingCardSpendingControlsAllowedCategoryDrinkingPlaces                                             IssuingCardSpendingControlsAllowedCategory = "drinking_places"
	IssuingCardSpendingControlsAllowedCategoryDrugStoresAndPharmacies                                    IssuingCardSpendingControlsAllowedCategory = "drug_stores_and_pharmacies"
	IssuingCardSpendingControlsAllowedCategoryDrugsDrugProprietariesAndDruggistSundries                  IssuingCardSpendingControlsAllowedCategory = "drugs_drug_proprietaries_and_druggist_sundries"
	IssuingCardSpendingControlsAllowedCategoryDryCleaners                                                IssuingCardSpendingControlsAllowedCategory = "dry_cleaners"
	IssuingCardSpendingControlsAllowedCategoryDurableGoods                                               IssuingCardSpendingControlsAllowedCategory = "durable_goods"
	IssuingCardSpendingControlsAllowedCategoryDutyFreeStores                                             IssuingCardSpendingControlsAllowedCategory = "duty_free_stores"
	IssuingCardSpendingControlsAllowedCategoryEatingPlacesRestaurants                                    IssuingCardSpendingControlsAllowedCategory = "eating_places_restaurants"
	IssuingCardSpendingControlsAllowedCategoryEducationalServices                                        IssuingCardSpendingControlsAllowedCategory = "educational_services"
	IssuingCardSpendingControlsAllowedCategoryElectricRazorStores                                        IssuingCardSpendingControlsAllowedCategory = "electric_razor_stores"
	IssuingCardSpendingControlsAllowedCategoryElectricalPartsAndEquipment                                IssuingCardSpendingControlsAllowedCategory = "electrical_parts_and_equipment"
	IssuingCardSpendingControlsAllowedCategoryElectricalServices                                         IssuingCardSpendingControlsAllowedCategory = "electrical_services"
	IssuingCardSpendingControlsAllowedCategoryElectronicsRepairShops                                     IssuingCardSpendingControlsAllowedCategory = "electronics_repair_shops"
	IssuingCardSpendingControlsAllowedCategoryElectronicsStores                                          IssuingCardSpendingControlsAllowedCategory = "electronics_stores"
	IssuingCardSpendingControlsAllowedCategoryElementarySecondarySchools                                 IssuingCardSpendingControlsAllowedCategory = "elementary_secondary_schools"
	IssuingCardSpendingControlsAllowedCategoryEmploymentTempAgencies                                     IssuingCardSpendingControlsAllowedCategory = "employment_temp_agencies"
	IssuingCardSpendingControlsAllowedCategoryEquipmentRental                                            IssuingCardSpendingControlsAllowedCategory = "equipment_rental"
	IssuingCardSpendingControlsAllowedCategoryExterminatingServices                                      IssuingCardSpendingControlsAllowedCategory = "exterminating_services"
	IssuingCardSpendingControlsAllowedCategoryFamilyClothingStores                                       IssuingCardSpendingControlsAllowedCategory = "family_clothing_stores"
	IssuingCardSpendingControlsAllowedCategoryFastFoodRestaurants                                        IssuingCardSpendingControlsAllowedCategory = "fast_food_restaurants"
	IssuingCardSpendingControlsAllowedCategoryFinancialInstitutions                                      IssuingCardSpendingControlsAllowedCategory = "financial_institutions"
	IssuingCardSpendingControlsAllowedCategoryFinesGovernmentAdministrativeEntities                      IssuingCardSpendingControlsAllowedCategory = "fines_government_administrative_entities"
	IssuingCardSpendingControlsAllowedCategoryFireplaceFireplaceScreensAndAccessoriesStores              IssuingCardSpendingControlsAllowedCategory = "fireplace_fireplace_screens_and_accessories_stores"
	IssuingCardSpendingControlsAllowedCategoryFloorCoveringStores                                        IssuingCardSpendingControlsAllowedCategory = "floor_covering_stores"
	IssuingCardSpendingControlsAllowedCategoryFlorists                                                   IssuingCardSpendingControlsAllowedCategory = "florists"
	IssuingCardSpendingControlsAllowedCategoryFloristsSuppliesNurseryStockAndFlowers                     IssuingCardSpendingControlsAllowedCategory = "florists_supplies_nursery_stock_and_flowers"
	IssuingCardSpendingControlsAllowedCategoryFreezerAndLockerMeatProvisioners                           IssuingCardSpendingControlsAllowedCategory = "freezer_and_locker_meat_provisioners"
	IssuingCardSpendingControlsAllowedCategoryFuelDealersNonAutomotive                                   IssuingCardSpendingControlsAllowedCategory = "fuel_dealers_non_automotive"
	IssuingCardSpendingControlsAllowedCategoryFuneralServicesCrematories                                 IssuingCardSpendingControlsAllowedCategory = "funeral_services_crematories"
	IssuingCardSpendingControlsAllowedCategoryFurnitureHomeFurnishingsAndEquipmentStoresExceptAppliances IssuingCardSpendingControlsAllowedCategory = "furniture_home_furnishings_and_equipment_stores_except_appliances"
	IssuingCardSpendingControlsAllowedCategoryFurnitureRepairRefinishing                                 IssuingCardSpendingControlsAllowedCategory = "furniture_repair_refinishing"
	IssuingCardSpendingControlsAllowedCategoryFurriersAndFurShops                                        IssuingCardSpendingControlsAllowedCategory = "furriers_and_fur_shops"
	IssuingCardSpendingControlsAllowedCategoryGeneralServices                                            IssuingCardSpendingControlsAllowedCategory = "general_services"
	IssuingCardSpendingControlsAllowedCategoryGiftCardNoveltyAndSouvenirShops                            IssuingCardSpendingControlsAllowedCategory = "gift_card_novelty_and_souvenir_shops"
	IssuingCardSpendingControlsAllowedCategoryGlassPaintAndWallpaperStores                               IssuingCardSpendingControlsAllowedCategory = "glass_paint_and_wallpaper_stores"
	IssuingCardSpendingControlsAllowedCategoryGlasswareCrystalStores                                     IssuingCardSpendingControlsAllowedCategory = "glassware_crystal_stores"
	IssuingCardSpendingControlsAllowedCategoryGolfCoursesPublic                                          IssuingCardSpendingControlsAllowedCategory = "golf_courses_public"
	IssuingCardSpendingControlsAllowedCategoryGovernmentServices                                         IssuingCardSpendingControlsAllowedCategory = "government_services"
	IssuingCardSpendingControlsAllowedCategoryGroceryStoresSupermarkets                                  IssuingCardSpendingControlsAllowedCategory = "grocery_stores_supermarkets"
	IssuingCardSpendingControlsAllowedCategoryHardwareEquipmentAndSupplies                               IssuingCardSpendingControlsAllowedCategory = "hardware_equipment_and_supplies"
	IssuingCardSpendingControlsAllowedCategoryHardwareStores                                             IssuingCardSpendingControlsAllowedCategory = "hardware_stores"
	IssuingCardSpendingControlsAllowedCategoryHealthAndBeautySpas                                        IssuingCardSpendingControlsAllowedCategory = "health_and_beauty_spas"
	IssuingCardSpendingControlsAllowedCategoryHearingAidsSalesAndSupplies                                IssuingCardSpendingControlsAllowedCategory = "hearing_aids_sales_and_supplies"
	IssuingCardSpendingControlsAllowedCategoryHeatingPlumbingAC                                          IssuingCardSpendingControlsAllowedCategory = "heating_plumbing_a_c"
	IssuingCardSpendingControlsAllowedCategoryHobbyToyAndGameShops                                       IssuingCardSpendingControlsAllowedCategory = "hobby_toy_and_game_shops"
	IssuingCardSpendingControlsAllowedCategoryHomeSupplyWarehouseStores                                  IssuingCardSpendingControlsAllowedCategory = "home_supply_warehouse_stores"
	IssuingCardSpendingControlsAllowedCategoryHospitals                                                  IssuingCardSpendingControlsAllowedCategory = "hospitals"
	IssuingCardSpendingControlsAllowedCategoryHotelsMotelsAndResorts                                     IssuingCardSpendingControlsAllowedCategory = "hotels_motels_and_resorts"
	IssuingCardSpendingControlsAllowedCategoryHouseholdApplianceStores                                   IssuingCardSpendingControlsAllowedCategory = "household_appliance_stores"
	IssuingCardSpendingControlsAllowedCategoryIndustrialSupplies                                         IssuingCardSpendingControlsAllowedCategory = "industrial_supplies"
	IssuingCardSpendingControlsAllowedCategoryInformationRetrievalServices                               IssuingCardSpendingControlsAllowedCategory = "information_retrieval_services"
	IssuingCardSpendingControlsAllowedCategoryInsuranceDefault                                           IssuingCardSpendingControlsAllowedCategory = "insurance_default"
	IssuingCardSpendingControlsAllowedCategoryInsuranceUnderwritingPremiums                              IssuingCardSpendingControlsAllowedCategory = "insurance_underwriting_premiums"
	IssuingCardSpendingControlsAllowedCategoryIntraCompanyPurchases                                      IssuingCardSpendingControlsAllowedCategory = "intra_company_purchases"
	IssuingCardSpendingControlsAllowedCategoryJewelryStoresWatchesClocksAndSilverwareStores              IssuingCardSpendingControlsAllowedCategory = "jewelry_stores_watches_clocks_and_silverware_stores"
	IssuingCardSpendingControlsAllowedCategoryLandscapingServices                                        IssuingCardSpendingControlsAllowedCategory = "landscaping_services"
	IssuingCardSpendingControlsAllowedCategoryLaundries                                                  IssuingCardSpendingControlsAllowedCategory = "laundries"
	IssuingCardSpendingControlsAllowedCategoryLaundryCleaningServices                                    IssuingCardSpendingControlsAllowedCategory = "laundry_cleaning_services"
	IssuingCardSpendingControlsAllowedCategoryLegalServicesAttorneys                                     IssuingCardSpendingControlsAllowedCategory = "legal_services_attorneys"
	IssuingCardSpendingControlsAllowedCategoryLuggageAndLeatherGoodsStores                               IssuingCardSpendingControlsAllowedCategory = "luggage_and_leather_goods_stores"
	IssuingCardSpendingControlsAllowedCategoryLumberBuildingMaterialsStores                              IssuingCardSpendingControlsAllowedCategory = "lumber_building_materials_stores"
	IssuingCardSpendingControlsAllowedCategoryManualCashDisburse                                         IssuingCardSpendingControlsAllowedCategory = "manual_cash_disburse"
	IssuingCardSpendingControlsAllowedCategoryMarinasServiceAndSupplies                                  IssuingCardSpendingControlsAllowedCategory = "marinas_service_and_supplies"
	IssuingCardSpendingControlsAllowedCategoryMasonryStoneworkAndPlaster                                 IssuingCardSpendingControlsAllowedCategory = "masonry_stonework_and_plaster"
	IssuingCardSpendingControlsAllowedCategoryMassageParlors                                             IssuingCardSpendingControlsAllowedCategory = "massage_parlors"
	IssuingCardSpendingControlsAllowedCategoryMedicalAndDentalLabs                                       IssuingCardSpendingControlsAllowedCategory = "medical_and_dental_labs"
	IssuingCardSpendingControlsAllowedCategoryMedicalDentalOphthalmicAndHospitalEquipmentAndSupplies     IssuingCardSpendingControlsAllowedCategory = "medical_dental_ophthalmic_and_hospital_equipment_and_supplies"
	IssuingCardSpendingControlsAllowedCategoryMedicalServices                                            IssuingCardSpendingControlsAllowedCategory = "medical_services"
	IssuingCardSpendingControlsAllowedCategoryMembershipOrganizations                                    IssuingCardSpendingControlsAllowedCategory = "membership_organizations"
	IssuingCardSpendingControlsAllowedCategoryMensAndBoysClothingAndAccessoriesStores                    IssuingCardSpendingControlsAllowedCategory = "mens_and_boys_clothing_and_accessories_stores"
	IssuingCardSpendingControlsAllowedCategoryMensWomensClothingStores                                   IssuingCardSpendingControlsAllowedCategory = "mens_womens_clothing_stores"
	IssuingCardSpendingControlsAllowedCategoryMetalServiceCenters                                        IssuingCardSpendingControlsAllowedCategory = "metal_service_centers"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneous                                              IssuingCardSpendingControlsAllowedCategory = "miscellaneous"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousApparelAndAccessoryShops                      IssuingCardSpendingControlsAllowedCategory = "miscellaneous_apparel_and_accessory_shops"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousAutoDealers                                   IssuingCardSpendingControlsAllowedCategory = "miscellaneous_auto_dealers"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousBusinessServices                              IssuingCardSpendingControlsAllowedCategory = "miscellaneous_business_services"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousFoodStores                                    IssuingCardSpendingControlsAllowedCategory = "miscellaneous_food_stores"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousGeneralMerchandise                            IssuingCardSpendingControlsAllowedCategory = "miscellaneous_general_merchandise"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousGeneralServices                               IssuingCardSpendingControlsAllowedCategory = "miscellaneous_general_services"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousHomeFurnishingSpecialtyStores                 IssuingCardSpendingControlsAllowedCategory = "miscellaneous_home_furnishing_specialty_stores"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousPublishingAndPrinting                         IssuingCardSpendingControlsAllowedCategory = "miscellaneous_publishing_and_printing"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousRecreationServices                            IssuingCardSpendingControlsAllowedCategory = "miscellaneous_recreation_services"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousRepairShops                                   IssuingCardSpendingControlsAllowedCategory = "miscellaneous_repair_shops"
	IssuingCardSpendingControlsAllowedCategoryMiscellaneousSpecialtyRetail                               IssuingCardSpendingControlsAllowedCategory = "miscellaneous_specialty_retail"
	IssuingCardSpendingControlsAllowedCategoryMobileHomeDealers                                          IssuingCardSpendingControlsAllowedCategory = "mobile_home_dealers"
	IssuingCardSpendingControlsAllowedCategoryMotionPictureTheaters                                      IssuingCardSpendingControlsAllowedCategory = "motion_picture_theaters"
	IssuingCardSpendingControlsAllowedCategoryMotorFreightCarriersAndTrucking                            IssuingCardSpendingControlsAllowedCategory = "motor_freight_carriers_and_trucking"
	IssuingCardSpendingControlsAllowedCategoryMotorHomesDealers                                          IssuingCardSpendingControlsAllowedCategory = "motor_homes_dealers"
	IssuingCardSpendingControlsAllowedCategoryMotorVehicleSuppliesAndNewParts                            IssuingCardSpendingControlsAllowedCategory = "motor_vehicle_supplies_and_new_parts"
	IssuingCardSpendingControlsAllowedCategoryMotorcycleShopsAndDealers                                  IssuingCardSpendingControlsAllowedCategory = "motorcycle_shops_and_dealers"
	IssuingCardSpendingControlsAllowedCategoryMotorcycleShopsDealers                                     IssuingCardSpendingControlsAllowedCategory = "motorcycle_shops_dealers"
	IssuingCardSpendingControlsAllowedCategoryMusicStoresMusicalInstrumentsPianosAndSheetMusic           IssuingCardSpendingControlsAllowedCategory = "music_stores_musical_instruments_pianos_and_sheet_music"
	IssuingCardSpendingControlsAllowedCategoryNewsDealersAndNewsstands                                   IssuingCardSpendingControlsAllowedCategory = "news_dealers_and_newsstands"
	IssuingCardSpendingControlsAllowedCategoryNonFiMoneyOrders                                           IssuingCardSpendingControlsAllowedCategory = "non_fi_money_orders"
	IssuingCardSpendingControlsAllowedCategoryNonFiStoredValueCardPurchaseLoad                           IssuingCardSpendingControlsAllowedCategory = "non_fi_stored_value_card_purchase_load"
	IssuingCardSpendingControlsAllowedCategoryNondurableGoods                                            IssuingCardSpendingControlsAllowedCategory = "nondurable_goods"
	IssuingCardSpendingControlsAllowedCategoryNurseriesLawnAndGardenSupplyStores                         IssuingCardSpendingControlsAllowedCategory = "nurseries_lawn_and_garden_supply_stores"
	IssuingCardSpendingControlsAllowedCategoryNursingPersonalCare                                        IssuingCardSpendingControlsAllowedCategory = "nursing_personal_care"
	IssuingCardSpendingControlsAllowedCategoryOfficeAndCommercialFurniture                               IssuingCardSpendingControlsAllowedCategory = "office_and_commercial_furniture"
	IssuingCardSpendingControlsAllowedCategoryOpticiansEyeglasses                                        IssuingCardSpendingControlsAllowedCategory = "opticians_eyeglasses"
	IssuingCardSpendingControlsAllowedCategoryOptometristsOphthalmologist                                IssuingCardSpendingControlsAllowedCategory = "optometrists_ophthalmologist"
	IssuingCardSpendingControlsAllowedCategoryOrthopedicGoodsProstheticDevices                           IssuingCardSpendingControlsAllowedCategory = "orthopedic_goods_prosthetic_devices"
	IssuingCardSpendingControlsAllowedCategoryOsteopaths                                                 IssuingCardSpendingControlsAllowedCategory = "osteopaths"
	IssuingCardSpendingControlsAllowedCategoryPackageStoresBeerWineAndLiquor                             IssuingCardSpendingControlsAllowedCategory = "package_stores_beer_wine_and_liquor"
	IssuingCardSpendingControlsAllowedCategoryPaintsVarnishesAndSupplies                                 IssuingCardSpendingControlsAllowedCategory = "paints_varnishes_and_supplies"
	IssuingCardSpendingControlsAllowedCategoryParkingLotsGarages                                         IssuingCardSpendingControlsAllowedCategory = "parking_lots_garages"
	IssuingCardSpendingControlsAllowedCategoryPassengerRailways                                          IssuingCardSpendingControlsAllowedCategory = "passenger_railways"
	IssuingCardSpendingControlsAllowedCategoryPawnShops                                                  IssuingCardSpendingControlsAllowedCategory = "pawn_shops"
	IssuingCardSpendingControlsAllowedCategoryPetShopsPetFoodAndSupplies                                 IssuingCardSpendingControlsAllowedCategory = "pet_shops_pet_food_and_supplies"
	IssuingCardSpendingControlsAllowedCategoryPetroleumAndPetroleumProducts                              IssuingCardSpendingControlsAllowedCategory = "petroleum_and_petroleum_products"
	IssuingCardSpendingControlsAllowedCategoryPhotoDeveloping                                            IssuingCardSpendingControlsAllowedCategory = "photo_developing"
	IssuingCardSpendingControlsAllowedCategoryPhotographicPhotocopyMicrofilmEquipmentAndSupplies         IssuingCardSpendingControlsAllowedCategory = "photographic_photocopy_microfilm_equipment_and_supplies"
	IssuingCardSpendingControlsAllowedCategoryPhotographicStudios                                        IssuingCardSpendingControlsAllowedCategory = "photographic_studios"
	IssuingCardSpendingControlsAllowedCategoryPictureVideoProduction                                     IssuingCardSpendingControlsAllowedCategory = "picture_video_production"
	IssuingCardSpendingControlsAllowedCategoryPieceGoodsNotionsAndOtherDryGoods                          IssuingCardSpendingControlsAllowedCategory = "piece_goods_notions_and_other_dry_goods"
	IssuingCardSpendingControlsAllowedCategoryPlumbingHeatingEquipmentAndSupplies                        IssuingCardSpendingControlsAllowedCategory = "plumbing_heating_equipment_and_supplies"
	IssuingCardSpendingControlsAllowedCategoryPoliticalOrganizations                                     IssuingCardSpendingControlsAllowedCategory = "political_organizations"
	IssuingCardSpendingControlsAllowedCategoryPostalServicesGovernmentOnly                               IssuingCardSpendingControlsAllowedCategory = "postal_services_government_only"
	IssuingCardSpendingControlsAllowedCategoryPreciousStonesAndMetalsWatchesAndJewelry                   IssuingCardSpendingControlsAllowedCategory = "precious_stones_and_metals_watches_and_jewelry"
	IssuingCardSpendingControlsAllowedCategoryProfessionalServices                                       IssuingCardSpendingControlsAllowedCategory = "professional_services"
	IssuingCardSpendingControlsAllowedCategoryPublicWarehousingAndStorage                                IssuingCardSpendingControlsAllowedCategory = "public_warehousing_and_storage"
	IssuingCardSpendingControlsAllowedCategoryQuickCopyReproAndBlueprint                                 IssuingCardSpendingControlsAllowedCategory = "quick_copy_repro_and_blueprint"
	IssuingCardSpendingControlsAllowedCategoryRailroads                                                  IssuingCardSpendingControlsAllowedCategory = "railroads"
	IssuingCardSpendingControlsAllowedCategoryRealEstateAgentsAndManagersRentals                         IssuingCardSpendingControlsAllowedCategory = "real_estate_agents_and_managers_rentals"
	IssuingCardSpendingControlsAllowedCategoryRecordStores                                               IssuingCardSpendingControlsAllowedCategory = "record_stores"
	IssuingCardSpendingControlsAllowedCategoryRecreationalVehicleRentals                                 IssuingCardSpendingControlsAllowedCategory = "recreational_vehicle_rentals"
	IssuingCardSpendingControlsAllowedCategoryReligiousGoodsStores                                       IssuingCardSpendingControlsAllowedCategory = "religious_goods_stores"
	IssuingCardSpendingControlsAllowedCategoryReligiousOrganizations                                     IssuingCardSpendingControlsAllowedCategory = "religious_organizations"
	IssuingCardSpendingControlsAllowedCategoryRoofingSidingSheetMetal                                    IssuingCardSpendingControlsAllowedCategory = "roofing_siding_sheet_metal"
	IssuingCardSpendingControlsAllowedCategorySecretarialSupportServices                                 IssuingCardSpendingControlsAllowedCategory = "secretarial_support_services"
	IssuingCardSpendingControlsAllowedCategorySecurityBrokersDealers                                     IssuingCardSpendingControlsAllowedCategory = "security_brokers_dealers"
	IssuingCardSpendingControlsAllowedCategoryServiceStations                                            IssuingCardSpendingControlsAllowedCategory = "service_stations"
	IssuingCardSpendingControlsAllowedCategorySewingNeedleworkFabricAndPieceGoodsStores                  IssuingCardSpendingControlsAllowedCategory = "sewing_needlework_fabric_and_piece_goods_stores"
	IssuingCardSpendingControlsAllowedCategoryShoeRepairHatCleaning                                      IssuingCardSpendingControlsAllowedCategory = "shoe_repair_hat_cleaning"
	IssuingCardSpendingControlsAllowedCategoryShoeStores                                                 IssuingCardSpendingControlsAllowedCategory = "shoe_stores"
	IssuingCardSpendingControlsAllowedCategorySmallApplianceRepair                                       IssuingCardSpendingControlsAllowedCategory = "small_appliance_repair"
	IssuingCardSpendingControlsAllowedCategorySnowmobileDealers                                          IssuingCardSpendingControlsAllowedCategory = "snowmobile_dealers"
	IssuingCardSpendingControlsAllowedCategorySpecialTradeServices                                       IssuingCardSpendingControlsAllowedCategory = "special_trade_services"
	IssuingCardSpendingControlsAllowedCategorySpecialtyCleaning                                          IssuingCardSpendingControlsAllowedCategory = "specialty_cleaning"
	IssuingCardSpendingControlsAllowedCategorySportingGoodsStores                                        IssuingCardSpendingControlsAllowedCategory = "sporting_goods_stores"
	IssuingCardSpendingControlsAllowedCategorySportingRecreationCamps                                    IssuingCardSpendingControlsAllowedCategory = "sporting_recreation_camps"
	IssuingCardSpendingControlsAllowedCategorySportsAndRidingApparelStores                               IssuingCardSpendingControlsAllowedCategory = "sports_and_riding_apparel_stores"
	IssuingCardSpendingControlsAllowedCategorySportsClubsFields                                          IssuingCardSpendingControlsAllowedCategory = "sports_clubs_fields"
	IssuingCardSpendingControlsAllowedCategoryStampAndCoinStores                                         IssuingCardSpendingControlsAllowedCategory = "stamp_and_coin_stores"
	IssuingCardSpendingControlsAllowedCategoryStationaryOfficeSuppliesPrintingAndWritingPaper            IssuingCardSpendingControlsAllowedCategory = "stationary_office_supplies_printing_and_writing_paper"
	IssuingCardSpendingControlsAllowedCategoryStationeryStoresOfficeAndSchoolSupplyStores                IssuingCardSpendingControlsAllowedCategory = "stationery_stores_office_and_school_supply_stores"
	IssuingCardSpendingControlsAllowedCategorySwimmingPoolsSales                                         IssuingCardSpendingControlsAllowedCategory = "swimming_pools_sales"
	IssuingCardSpendingControlsAllowedCategoryTUiTravelGermany                                           IssuingCardSpendingControlsAllowedCategory = "t_ui_travel_germany"
	IssuingCardSpendingControlsAllowedCategoryTailorsAlterations                                         IssuingCardSpendingControlsAllowedCategory = "tailors_alterations"
	IssuingCardSpendingControlsAllowedCategoryTaxPaymentsGovernmentAgencies                              IssuingCardSpendingControlsAllowedCategory = "tax_payments_government_agencies"
	IssuingCardSpendingControlsAllowedCategoryTaxPreparationServices                                     IssuingCardSpendingControlsAllowedCategory = "tax_preparation_services"
	IssuingCardSpendingControlsAllowedCategoryTaxicabsLimousines                                         IssuingCardSpendingControlsAllowedCategory = "taxicabs_limousines"
	IssuingCardSpendingControlsAllowedCategoryTelecommunicationEquipmentAndTelephoneSales                IssuingCardSpendingControlsAllowedCategory = "telecommunication_equipment_and_telephone_sales"
	IssuingCardSpendingControlsAllowedCategoryTelecommunicationServices                                  IssuingCardSpendingControlsAllowedCategory = "telecommunication_services"
	IssuingCardSpendingControlsAllowedCategoryTelegraphServices                                          IssuingCardSpendingControlsAllowedCategory = "telegraph_services"
	IssuingCardSpendingControlsAllowedCategoryTentAndAwningShops                                         IssuingCardSpendingControlsAllowedCategory = "tent_and_awning_shops"
	IssuingCardSpendingControlsAllowedCategoryTestingLaboratories                                        IssuingCardSpendingControlsAllowedCategory = "testing_laboratories"
	IssuingCardSpendingControlsAllowedCategoryTheatricalTicketAgencies                                   IssuingCardSpendingControlsAllowedCategory = "theatrical_ticket_agencies"
	IssuingCardSpendingControlsAllowedCategoryTimeshares                                                 IssuingCardSpendingControlsAllowedCategory = "timeshares"
	IssuingCardSpendingControlsAllowedCategoryTireRetreadingAndRepair                                    IssuingCardSpendingControlsAllowedCategory = "tire_retreading_and_repair"
	IssuingCardSpendingControlsAllowedCategoryTollsBridgeFees                                            IssuingCardSpendingControlsAllowedCategory = "tolls_bridge_fees"
	IssuingCardSpendingControlsAllowedCategoryTouristAttractionsAndExhibits                              IssuingCardSpendingControlsAllowedCategory = "tourist_attractions_and_exhibits"
	IssuingCardSpendingControlsAllowedCategoryTowingServices                                             IssuingCardSpendingControlsAllowedCategory = "towing_services"
	IssuingCardSpendingControlsAllowedCategoryTrailerParksCampgrounds                                    IssuingCardSpendingControlsAllowedCategory = "trailer_parks_campgrounds"
	IssuingCardSpendingControlsAllowedCategoryTransportationServices                                     IssuingCardSpendingControlsAllowedCategory = "transportation_services"
	IssuingCardSpendingControlsAllowedCategoryTravelAgenciesTourOperators                                IssuingCardSpendingControlsAllowedCategory = "travel_agencies_tour_operators"
	IssuingCardSpendingControlsAllowedCategoryTruckStopIteration                                         IssuingCardSpendingControlsAllowedCategory = "truck_stop_iteration"
	IssuingCardSpendingControlsAllowedCategoryTruckUtilityTrailerRentals                                 IssuingCardSpendingControlsAllowedCategory = "truck_utility_trailer_rentals"
	IssuingCardSpendingControlsAllowedCategoryTypesettingPlateMakingAndRelatedServices                   IssuingCardSpendingControlsAllowedCategory = "typesetting_plate_making_and_related_services"
	IssuingCardSpendingControlsAllowedCategoryTypewriterStores                                           IssuingCardSpendingControlsAllowedCategory = "typewriter_stores"
	IssuingCardSpendingControlsAllowedCategoryUSFederalGovernmentAgenciesOrDepartments                   IssuingCardSpendingControlsAllowedCategory = "u_s_federal_government_agencies_or_departments"
	IssuingCardSpendingControlsAllowedCategoryUniformsCommercialClothing                                 IssuingCardSpendingControlsAllowedCategory = "uniforms_commercial_clothing"
	IssuingCardSpendingControlsAllowedCategoryUsedMerchandiseAndSecondhandStores                         IssuingCardSpendingControlsAllowedCategory = "used_merchandise_and_secondhand_stores"
	IssuingCardSpendingControlsAllowedCategoryUtilities                                                  IssuingCardSpendingControlsAllowedCategory = "utilities"
	IssuingCardSpendingControlsAllowedCategoryVarietyStores                                              IssuingCardSpendingControlsAllowedCategory = "variety_stores"
	IssuingCardSpendingControlsAllowedCategoryVeterinaryServices                                         IssuingCardSpendingControlsAllowedCategory = "veterinary_services"
	IssuingCardSpendingControlsAllowedCategoryVideoAmusementGameSupplies                                 IssuingCardSpendingControlsAllowedCategory = "video_amusement_game_supplies"
	IssuingCardSpendingControlsAllowedCategoryVideoGameArcades                                           IssuingCardSpendingControlsAllowedCategory = "video_game_arcades"
	IssuingCardSpendingControlsAllowedCategoryVideoTapeRentalStores                                      IssuingCardSpendingControlsAllowedCategory = "video_tape_rental_stores"
	IssuingCardSpendingControlsAllowedCategoryVocationalTradeSchools                                     IssuingCardSpendingControlsAllowedCategory = "vocational_trade_schools"
	IssuingCardSpendingControlsAllowedCategoryWatchJewelryRepair                                         IssuingCardSpendingControlsAllowedCategory = "watch_jewelry_repair"
	IssuingCardSpendingControlsAllowedCategoryWeldingRepair                                              IssuingCardSpendingControlsAllowedCategory = "welding_repair"
	IssuingCardSpendingControlsAllowedCategoryWholesaleClubs                                             IssuingCardSpendingControlsAllowedCategory = "wholesale_clubs"
	IssuingCardSpendingControlsAllowedCategoryWigAndToupeeStores                                         IssuingCardSpendingControlsAllowedCategory = "wig_and_toupee_stores"
	IssuingCardSpendingControlsAllowedCategoryWiresMoneyOrders                                           IssuingCardSpendingControlsAllowedCategory = "wires_money_orders"
	IssuingCardSpendingControlsAllowedCategoryWomensAccessoryAndSpecialtyShops                           IssuingCardSpendingControlsAllowedCategory = "womens_accessory_and_specialty_shops"
	IssuingCardSpendingControlsAllowedCategoryWomensReadyToWearStores                                    IssuingCardSpendingControlsAllowedCategory = "womens_ready_to_wear_stores"
	IssuingCardSpendingControlsAllowedCategoryWreckingAndSalvageYards                                    IssuingCardSpendingControlsAllowedCategory = "wrecking_and_salvage_yards"
)

// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) of authorizations to decline. All other categories will be allowed. Cannot be set with `allowed_categories`.
type IssuingCardSpendingControlsBlockedCategory string

// List of values that IssuingCardSpendingControlsBlockedCategory can take
const (
	IssuingCardSpendingControlsBlockedCategoryAcRefrigerationRepair                                      IssuingCardSpendingControlsBlockedCategory = "ac_refrigeration_repair"
	IssuingCardSpendingControlsBlockedCategoryAccountingBookkeepingServices                              IssuingCardSpendingControlsBlockedCategory = "accounting_bookkeeping_services"
	IssuingCardSpendingControlsBlockedCategoryAdvertisingServices                                        IssuingCardSpendingControlsBlockedCategory = "advertising_services"
	IssuingCardSpendingControlsBlockedCategoryAgriculturalCooperative                                    IssuingCardSpendingControlsBlockedCategory = "agricultural_cooperative"
	IssuingCardSpendingControlsBlockedCategoryAirlinesAirCarriers                                        IssuingCardSpendingControlsBlockedCategory = "airlines_air_carriers"
	IssuingCardSpendingControlsBlockedCategoryAirportsFlyingFields                                       IssuingCardSpendingControlsBlockedCategory = "airports_flying_fields"
	IssuingCardSpendingControlsBlockedCategoryAmbulanceServices                                          IssuingCardSpendingControlsBlockedCategory = "ambulance_services"
	IssuingCardSpendingControlsBlockedCategoryAmusementParksCarnivals                                    IssuingCardSpendingControlsBlockedCategory = "amusement_parks_carnivals"
	IssuingCardSpendingControlsBlockedCategoryAntiqueReproductions                                       IssuingCardSpendingControlsBlockedCategory = "antique_reproductions"
	IssuingCardSpendingControlsBlockedCategoryAntiqueShops                                               IssuingCardSpendingControlsBlockedCategory = "antique_shops"
	IssuingCardSpendingControlsBlockedCategoryAquariums                                                  IssuingCardSpendingControlsBlockedCategory = "aquariums"
	IssuingCardSpendingControlsBlockedCategoryArchitecturalSurveyingServices                             IssuingCardSpendingControlsBlockedCategory = "architectural_surveying_services"
	IssuingCardSpendingControlsBlockedCategoryArtDealersAndGalleries                                     IssuingCardSpendingControlsBlockedCategory = "art_dealers_and_galleries"
	IssuingCardSpendingControlsBlockedCategoryArtistsSupplyAndCraftShops                                 IssuingCardSpendingControlsBlockedCategory = "artists_supply_and_craft_shops"
	IssuingCardSpendingControlsBlockedCategoryAutoAndHomeSupplyStores                                    IssuingCardSpendingControlsBlockedCategory = "auto_and_home_supply_stores"
	IssuingCardSpendingControlsBlockedCategoryAutoBodyRepairShops                                        IssuingCardSpendingControlsBlockedCategory = "auto_body_repair_shops"
	IssuingCardSpendingControlsBlockedCategoryAutoPaintShops                                             IssuingCardSpendingControlsBlockedCategory = "auto_paint_shops"
	IssuingCardSpendingControlsBlockedCategoryAutoServiceShops                                           IssuingCardSpendingControlsBlockedCategory = "auto_service_shops"
	IssuingCardSpendingControlsBlockedCategoryAutomatedCashDisburse                                      IssuingCardSpendingControlsBlockedCategory = "automated_cash_disburse"
	IssuingCardSpendingControlsBlockedCategoryAutomatedFuelDispensers                                    IssuingCardSpendingControlsBlockedCategory = "automated_fuel_dispensers"
	IssuingCardSpendingControlsBlockedCategoryAutomobileAssociations                                     IssuingCardSpendingControlsBlockedCategory = "automobile_associations"
	IssuingCardSpendingControlsBlockedCategoryAutomotivePartsAndAccessoriesStores                        IssuingCardSpendingControlsBlockedCategory = "automotive_parts_and_accessories_stores"
	IssuingCardSpendingControlsBlockedCategoryAutomotiveTireStores                                       IssuingCardSpendingControlsBlockedCategory = "automotive_tire_stores"
	IssuingCardSpendingControlsBlockedCategoryBailAndBondPayments                                        IssuingCardSpendingControlsBlockedCategory = "bail_and_bond_payments"
	IssuingCardSpendingControlsBlockedCategoryBakeries                                                   IssuingCardSpendingControlsBlockedCategory = "bakeries"
	IssuingCardSpendingControlsBlockedCategoryBandsOrchestras                                            IssuingCardSpendingControlsBlockedCategory = "bands_orchestras"
	IssuingCardSpendingControlsBlockedCategoryBarberAndBeautyShops                                       IssuingCardSpendingControlsBlockedCategory = "barber_and_beauty_shops"
	IssuingCardSpendingControlsBlockedCategoryBettingCasinoGambling                                      IssuingCardSpendingControlsBlockedCategory = "betting_casino_gambling"
	IssuingCardSpendingControlsBlockedCategoryBicycleShops                                               IssuingCardSpendingControlsBlockedCategory = "bicycle_shops"
	IssuingCardSpendingControlsBlockedCategoryBilliardPoolEstablishments                                 IssuingCardSpendingControlsBlockedCategory = "billiard_pool_establishments"
	IssuingCardSpendingControlsBlockedCategoryBoatDealers                                                IssuingCardSpendingControlsBlockedCategory = "boat_dealers"
	IssuingCardSpendingControlsBlockedCategoryBoatRentalsAndLeases                                       IssuingCardSpendingControlsBlockedCategory = "boat_rentals_and_leases"
	IssuingCardSpendingControlsBlockedCategoryBookStores                                                 IssuingCardSpendingControlsBlockedCategory = "book_stores"
	IssuingCardSpendingControlsBlockedCategoryBooksPeriodicalsAndNewspapers                              IssuingCardSpendingControlsBlockedCategory = "books_periodicals_and_newspapers"
	IssuingCardSpendingControlsBlockedCategoryBowlingAlleys                                              IssuingCardSpendingControlsBlockedCategory = "bowling_alleys"
	IssuingCardSpendingControlsBlockedCategoryBusLines                                                   IssuingCardSpendingControlsBlockedCategory = "bus_lines"
	IssuingCardSpendingControlsBlockedCategoryBusinessSecretarialSchools                                 IssuingCardSpendingControlsBlockedCategory = "business_secretarial_schools"
	IssuingCardSpendingControlsBlockedCategoryBuyingShoppingServices                                     IssuingCardSpendingControlsBlockedCategory = "buying_shopping_services"
	IssuingCardSpendingControlsBlockedCategoryCableSatelliteAndOtherPayTelevisionAndRadio                IssuingCardSpendingControlsBlockedCategory = "cable_satellite_and_other_pay_television_and_radio"
	IssuingCardSpendingControlsBlockedCategoryCameraAndPhotographicSupplyStores                          IssuingCardSpendingControlsBlockedCategory = "camera_and_photographic_supply_stores"
	IssuingCardSpendingControlsBlockedCategoryCandyNutAndConfectioneryStores                             IssuingCardSpendingControlsBlockedCategory = "candy_nut_and_confectionery_stores"
	IssuingCardSpendingControlsBlockedCategoryCarAndTruckDealersNewUsed                                  IssuingCardSpendingControlsBlockedCategory = "car_and_truck_dealers_new_used"
	IssuingCardSpendingControlsBlockedCategoryCarAndTruckDealersUsedOnly                                 IssuingCardSpendingControlsBlockedCategory = "car_and_truck_dealers_used_only"
	IssuingCardSpendingControlsBlockedCategoryCarRentalAgencies                                          IssuingCardSpendingControlsBlockedCategory = "car_rental_agencies"
	IssuingCardSpendingControlsBlockedCategoryCarWashes                                                  IssuingCardSpendingControlsBlockedCategory = "car_washes"
	IssuingCardSpendingControlsBlockedCategoryCarpentryServices                                          IssuingCardSpendingControlsBlockedCategory = "carpentry_services"
	IssuingCardSpendingControlsBlockedCategoryCarpetUpholsteryCleaning                                   IssuingCardSpendingControlsBlockedCategory = "carpet_upholstery_cleaning"
	IssuingCardSpendingControlsBlockedCategoryCaterers                                                   IssuingCardSpendingControlsBlockedCategory = "caterers"
	IssuingCardSpendingControlsBlockedCategoryCharitableAndSocialServiceOrganizationsFundraising         IssuingCardSpendingControlsBlockedCategory = "charitable_and_social_service_organizations_fundraising"
	IssuingCardSpendingControlsBlockedCategoryChemicalsAndAlliedProducts                                 IssuingCardSpendingControlsBlockedCategory = "chemicals_and_allied_products"
	IssuingCardSpendingControlsBlockedCategoryChildCareServices                                          IssuingCardSpendingControlsBlockedCategory = "child_care_services"
	IssuingCardSpendingControlsBlockedCategoryChildrensAndInfantsWearStores                              IssuingCardSpendingControlsBlockedCategory = "childrens_and_infants_wear_stores"
	IssuingCardSpendingControlsBlockedCategoryChiropodistsPodiatrists                                    IssuingCardSpendingControlsBlockedCategory = "chiropodists_podiatrists"
	IssuingCardSpendingControlsBlockedCategoryChiropractors                                              IssuingCardSpendingControlsBlockedCategory = "chiropractors"
	IssuingCardSpendingControlsBlockedCategoryCigarStoresAndStands                                       IssuingCardSpendingControlsBlockedCategory = "cigar_stores_and_stands"
	IssuingCardSpendingControlsBlockedCategoryCivicSocialFraternalAssociations                           IssuingCardSpendingControlsBlockedCategory = "civic_social_fraternal_associations"
	IssuingCardSpendingControlsBlockedCategoryCleaningAndMaintenance                                     IssuingCardSpendingControlsBlockedCategory = "cleaning_and_maintenance"
	IssuingCardSpendingControlsBlockedCategoryClothingRental                                             IssuingCardSpendingControlsBlockedCategory = "clothing_rental"
	IssuingCardSpendingControlsBlockedCategoryCollegesUniversities                                       IssuingCardSpendingControlsBlockedCategory = "colleges_universities"
	IssuingCardSpendingControlsBlockedCategoryCommercialEquipment                                        IssuingCardSpendingControlsBlockedCategory = "commercial_equipment"
	IssuingCardSpendingControlsBlockedCategoryCommercialFootwear                                         IssuingCardSpendingControlsBlockedCategory = "commercial_footwear"
	IssuingCardSpendingControlsBlockedCategoryCommercialPhotographyArtAndGraphics                        IssuingCardSpendingControlsBlockedCategory = "commercial_photography_art_and_graphics"
	IssuingCardSpendingControlsBlockedCategoryCommuterTransportAndFerries                                IssuingCardSpendingControlsBlockedCategory = "commuter_transport_and_ferries"
	IssuingCardSpendingControlsBlockedCategoryComputerNetworkServices                                    IssuingCardSpendingControlsBlockedCategory = "computer_network_services"
	IssuingCardSpendingControlsBlockedCategoryComputerProgramming                                        IssuingCardSpendingControlsBlockedCategory = "computer_programming"
	IssuingCardSpendingControlsBlockedCategoryComputerRepair                                             IssuingCardSpendingControlsBlockedCategory = "computer_repair"
	IssuingCardSpendingControlsBlockedCategoryComputerSoftwareStores                                     IssuingCardSpendingControlsBlockedCategory = "computer_software_stores"
	IssuingCardSpendingControlsBlockedCategoryComputersPeripheralsAndSoftware                            IssuingCardSpendingControlsBlockedCategory = "computers_peripherals_and_software"
	IssuingCardSpendingControlsBlockedCategoryConcreteWorkServices                                       IssuingCardSpendingControlsBlockedCategory = "concrete_work_services"
	IssuingCardSpendingControlsBlockedCategoryConstructionMaterials                                      IssuingCardSpendingControlsBlockedCategory = "construction_materials"
	IssuingCardSpendingControlsBlockedCategoryConsultingPublicRelations                                  IssuingCardSpendingControlsBlockedCategory = "consulting_public_relations"
	IssuingCardSpendingControlsBlockedCategoryCorrespondenceSchools                                      IssuingCardSpendingControlsBlockedCategory = "correspondence_schools"
	IssuingCardSpendingControlsBlockedCategoryCosmeticStores                                             IssuingCardSpendingControlsBlockedCategory = "cosmetic_stores"
	IssuingCardSpendingControlsBlockedCategoryCounselingServices                                         IssuingCardSpendingControlsBlockedCategory = "counseling_services"
	IssuingCardSpendingControlsBlockedCategoryCountryClubs                                               IssuingCardSpendingControlsBlockedCategory = "country_clubs"
	IssuingCardSpendingControlsBlockedCategoryCourierServices                                            IssuingCardSpendingControlsBlockedCategory = "courier_services"
	IssuingCardSpendingControlsBlockedCategoryCourtCosts                                                 IssuingCardSpendingControlsBlockedCategory = "court_costs"
	IssuingCardSpendingControlsBlockedCategoryCreditReportingAgencies                                    IssuingCardSpendingControlsBlockedCategory = "credit_reporting_agencies"
	IssuingCardSpendingControlsBlockedCategoryCruiseLines                                                IssuingCardSpendingControlsBlockedCategory = "cruise_lines"
	IssuingCardSpendingControlsBlockedCategoryDairyProductsStores                                        IssuingCardSpendingControlsBlockedCategory = "dairy_products_stores"
	IssuingCardSpendingControlsBlockedCategoryDanceHallStudiosSchools                                    IssuingCardSpendingControlsBlockedCategory = "dance_hall_studios_schools"
	IssuingCardSpendingControlsBlockedCategoryDatingEscortServices                                       IssuingCardSpendingControlsBlockedCategory = "dating_escort_services"
	IssuingCardSpendingControlsBlockedCategoryDentistsOrthodontists                                      IssuingCardSpendingControlsBlockedCategory = "dentists_orthodontists"
	IssuingCardSpendingControlsBlockedCategoryDepartmentStores                                           IssuingCardSpendingControlsBlockedCategory = "department_stores"
	IssuingCardSpendingControlsBlockedCategoryDetectiveAgencies                                          IssuingCardSpendingControlsBlockedCategory = "detective_agencies"
	IssuingCardSpendingControlsBlockedCategoryDigitalGoodsApplications                                   IssuingCardSpendingControlsBlockedCategory = "digital_goods_applications"
	IssuingCardSpendingControlsBlockedCategoryDigitalGoodsGames                                          IssuingCardSpendingControlsBlockedCategory = "digital_goods_games"
	IssuingCardSpendingControlsBlockedCategoryDigitalGoodsLargeVolume                                    IssuingCardSpendingControlsBlockedCategory = "digital_goods_large_volume"
	IssuingCardSpendingControlsBlockedCategoryDigitalGoodsMedia                                          IssuingCardSpendingControlsBlockedCategory = "digital_goods_media"
	IssuingCardSpendingControlsBlockedCategoryDirectMarketingCatalogMerchant                             IssuingCardSpendingControlsBlockedCategory = "direct_marketing_catalog_merchant"
	IssuingCardSpendingControlsBlockedCategoryDirectMarketingCombinationCatalogAndRetailMerchant         IssuingCardSpendingControlsBlockedCategory = "direct_marketing_combination_catalog_and_retail_merchant"
	IssuingCardSpendingControlsBlockedCategoryDirectMarketingInboundTelemarketing                        IssuingCardSpendingControlsBlockedCategory = "direct_marketing_inbound_telemarketing"
	IssuingCardSpendingControlsBlockedCategoryDirectMarketingInsuranceServices                           IssuingCardSpendingControlsBlockedCategory = "direct_marketing_insurance_services"
	IssuingCardSpendingControlsBlockedCategoryDirectMarketingOther                                       IssuingCardSpendingControlsBlockedCategory = "direct_marketing_other"
	IssuingCardSpendingControlsBlockedCategoryDirectMarketingOutboundTelemarketing                       IssuingCardSpendingControlsBlockedCategory = "direct_marketing_outbound_telemarketing"
	IssuingCardSpendingControlsBlockedCategoryDirectMarketingSubscription                                IssuingCardSpendingControlsBlockedCategory = "direct_marketing_subscription"
	IssuingCardSpendingControlsBlockedCategoryDirectMarketingTravel                                      IssuingCardSpendingControlsBlockedCategory = "direct_marketing_travel"
	IssuingCardSpendingControlsBlockedCategoryDiscountStores                                             IssuingCardSpendingControlsBlockedCategory = "discount_stores"
	IssuingCardSpendingControlsBlockedCategoryDoctors                                                    IssuingCardSpendingControlsBlockedCategory = "doctors"
	IssuingCardSpendingControlsBlockedCategoryDoorToDoorSales                                            IssuingCardSpendingControlsBlockedCategory = "door_to_door_sales"
	IssuingCardSpendingControlsBlockedCategoryDraperyWindowCoveringAndUpholsteryStores                   IssuingCardSpendingControlsBlockedCategory = "drapery_window_covering_and_upholstery_stores"
	IssuingCardSpendingControlsBlockedCategoryDrinkingPlaces                                             IssuingCardSpendingControlsBlockedCategory = "drinking_places"
	IssuingCardSpendingControlsBlockedCategoryDrugStoresAndPharmacies                                    IssuingCardSpendingControlsBlockedCategory = "drug_stores_and_pharmacies"
	IssuingCardSpendingControlsBlockedCategoryDrugsDrugProprietariesAndDruggistSundries                  IssuingCardSpendingControlsBlockedCategory = "drugs_drug_proprietaries_and_druggist_sundries"
	IssuingCardSpendingControlsBlockedCategoryDryCleaners                                                IssuingCardSpendingControlsBlockedCategory = "dry_cleaners"
	IssuingCardSpendingControlsBlockedCategoryDurableGoods                                               IssuingCardSpendingControlsBlockedCategory = "durable_goods"
	IssuingCardSpendingControlsBlockedCategoryDutyFreeStores                                             IssuingCardSpendingControlsBlockedCategory = "duty_free_stores"
	IssuingCardSpendingControlsBlockedCategoryEatingPlacesRestaurants                                    IssuingCardSpendingControlsBlockedCategory = "eating_places_restaurants"
	IssuingCardSpendingControlsBlockedCategoryEducationalServices                                        IssuingCardSpendingControlsBlockedCategory = "educational_services"
	IssuingCardSpendingControlsBlockedCategoryElectricRazorStores                                        IssuingCardSpendingControlsBlockedCategory = "electric_razor_stores"
	IssuingCardSpendingControlsBlockedCategoryElectricalPartsAndEquipment                                IssuingCardSpendingControlsBlockedCategory = "electrical_parts_and_equipment"
	IssuingCardSpendingControlsBlockedCategoryElectricalServices                                         IssuingCardSpendingControlsBlockedCategory = "electrical_services"
	IssuingCardSpendingControlsBlockedCategoryElectronicsRepairShops                                     IssuingCardSpendingControlsBlockedCategory = "electronics_repair_shops"
	IssuingCardSpendingControlsBlockedCategoryElectronicsStores                                          IssuingCardSpendingControlsBlockedCategory = "electronics_stores"
	IssuingCardSpendingControlsBlockedCategoryElementarySecondarySchools                                 IssuingCardSpendingControlsBlockedCategory = "elementary_secondary_schools"
	IssuingCardSpendingControlsBlockedCategoryEmploymentTempAgencies                                     IssuingCardSpendingControlsBlockedCategory = "employment_temp_agencies"
	IssuingCardSpendingControlsBlockedCategoryEquipmentRental                                            IssuingCardSpendingControlsBlockedCategory = "equipment_rental"
	IssuingCardSpendingControlsBlockedCategoryExterminatingServices                                      IssuingCardSpendingControlsBlockedCategory = "exterminating_services"
	IssuingCardSpendingControlsBlockedCategoryFamilyClothingStores                                       IssuingCardSpendingControlsBlockedCategory = "family_clothing_stores"
	IssuingCardSpendingControlsBlockedCategoryFastFoodRestaurants                                        IssuingCardSpendingControlsBlockedCategory = "fast_food_restaurants"
	IssuingCardSpendingControlsBlockedCategoryFinancialInstitutions                                      IssuingCardSpendingControlsBlockedCategory = "financial_institutions"
	IssuingCardSpendingControlsBlockedCategoryFinesGovernmentAdministrativeEntities                      IssuingCardSpendingControlsBlockedCategory = "fines_government_administrative_entities"
	IssuingCardSpendingControlsBlockedCategoryFireplaceFireplaceScreensAndAccessoriesStores              IssuingCardSpendingControlsBlockedCategory = "fireplace_fireplace_screens_and_accessories_stores"
	IssuingCardSpendingControlsBlockedCategoryFloorCoveringStores                                        IssuingCardSpendingControlsBlockedCategory = "floor_covering_stores"
	IssuingCardSpendingControlsBlockedCategoryFlorists                                                   IssuingCardSpendingControlsBlockedCategory = "florists"
	IssuingCardSpendingControlsBlockedCategoryFloristsSuppliesNurseryStockAndFlowers                     IssuingCardSpendingControlsBlockedCategory = "florists_supplies_nursery_stock_and_flowers"
	IssuingCardSpendingControlsBlockedCategoryFreezerAndLockerMeatProvisioners                           IssuingCardSpendingControlsBlockedCategory = "freezer_and_locker_meat_provisioners"
	IssuingCardSpendingControlsBlockedCategoryFuelDealersNonAutomotive                                   IssuingCardSpendingControlsBlockedCategory = "fuel_dealers_non_automotive"
	IssuingCardSpendingControlsBlockedCategoryFuneralServicesCrematories                                 IssuingCardSpendingControlsBlockedCategory = "funeral_services_crematories"
	IssuingCardSpendingControlsBlockedCategoryFurnitureHomeFurnishingsAndEquipmentStoresExceptAppliances IssuingCardSpendingControlsBlockedCategory = "furniture_home_furnishings_and_equipment_stores_except_appliances"
	IssuingCardSpendingControlsBlockedCategoryFurnitureRepairRefinishing                                 IssuingCardSpendingControlsBlockedCategory = "furniture_repair_refinishing"
	IssuingCardSpendingControlsBlockedCategoryFurriersAndFurShops                                        IssuingCardSpendingControlsBlockedCategory = "furriers_and_fur_shops"
	IssuingCardSpendingControlsBlockedCategoryGeneralServices                                            IssuingCardSpendingControlsBlockedCategory = "general_services"
	IssuingCardSpendingControlsBlockedCategoryGiftCardNoveltyAndSouvenirShops                            IssuingCardSpendingControlsBlockedCategory = "gift_card_novelty_and_souvenir_shops"
	IssuingCardSpendingControlsBlockedCategoryGlassPaintAndWallpaperStores                               IssuingCardSpendingControlsBlockedCategory = "glass_paint_and_wallpaper_stores"
	IssuingCardSpendingControlsBlockedCategoryGlasswareCrystalStores                                     IssuingCardSpendingControlsBlockedCategory = "glassware_crystal_stores"
	IssuingCardSpendingControlsBlockedCategoryGolfCoursesPublic                                          IssuingCardSpendingControlsBlockedCategory = "golf_courses_public"
	IssuingCardSpendingControlsBlockedCategoryGovernmentServices                                         IssuingCardSpendingControlsBlockedCategory = "government_services"
	IssuingCardSpendingControlsBlockedCategoryGroceryStoresSupermarkets                                  IssuingCardSpendingControlsBlockedCategory = "grocery_stores_supermarkets"
	IssuingCardSpendingControlsBlockedCategoryHardwareEquipmentAndSupplies                               IssuingCardSpendingControlsBlockedCategory = "hardware_equipment_and_supplies"
	IssuingCardSpendingControlsBlockedCategoryHardwareStores                                             IssuingCardSpendingControlsBlockedCategory = "hardware_stores"
	IssuingCardSpendingControlsBlockedCategoryHealthAndBeautySpas                                        IssuingCardSpendingControlsBlockedCategory = "health_and_beauty_spas"
	IssuingCardSpendingControlsBlockedCategoryHearingAidsSalesAndSupplies                                IssuingCardSpendingControlsBlockedCategory = "hearing_aids_sales_and_supplies"
	IssuingCardSpendingControlsBlockedCategoryHeatingPlumbingAC                                          IssuingCardSpendingControlsBlockedCategory = "heating_plumbing_a_c"
	IssuingCardSpendingControlsBlockedCategoryHobbyToyAndGameShops                                       IssuingCardSpendingControlsBlockedCategory = "hobby_toy_and_game_shops"
	IssuingCardSpendingControlsBlockedCategoryHomeSupplyWarehouseStores                                  IssuingCardSpendingControlsBlockedCategory = "home_supply_warehouse_stores"
	IssuingCardSpendingControlsBlockedCategoryHospitals                                                  IssuingCardSpendingControlsBlockedCategory = "hospitals"
	IssuingCardSpendingControlsBlockedCategoryHotelsMotelsAndResorts                                     IssuingCardSpendingControlsBlockedCategory = "hotels_motels_and_resorts"
	IssuingCardSpendingControlsBlockedCategoryHouseholdApplianceStores                                   IssuingCardSpendingControlsBlockedCategory = "household_appliance_stores"
	IssuingCardSpendingControlsBlockedCategoryIndustrialSupplies                                         IssuingCardSpendingControlsBlockedCategory = "industrial_supplies"
	IssuingCardSpendingControlsBlockedCategoryInformationRetrievalServices                               IssuingCardSpendingControlsBlockedCategory = "information_retrieval_services"
	IssuingCardSpendingControlsBlockedCategoryInsuranceDefault                                           IssuingCardSpendingControlsBlockedCategory = "insurance_default"
	IssuingCardSpendingControlsBlockedCategoryInsuranceUnderwritingPremiums                              IssuingCardSpendingControlsBlockedCategory = "insurance_underwriting_premiums"
	IssuingCardSpendingControlsBlockedCategoryIntraCompanyPurchases                                      IssuingCardSpendingControlsBlockedCategory = "intra_company_purchases"
	IssuingCardSpendingControlsBlockedCategoryJewelryStoresWatchesClocksAndSilverwareStores              IssuingCardSpendingControlsBlockedCategory = "jewelry_stores_watches_clocks_and_silverware_stores"
	IssuingCardSpendingControlsBlockedCategoryLandscapingServices                                        IssuingCardSpendingControlsBlockedCategory = "landscaping_services"
	IssuingCardSpendingControlsBlockedCategoryLaundries                                                  IssuingCardSpendingControlsBlockedCategory = "laundries"
	IssuingCardSpendingControlsBlockedCategoryLaundryCleaningServices                                    IssuingCardSpendingControlsBlockedCategory = "laundry_cleaning_services"
	IssuingCardSpendingControlsBlockedCategoryLegalServicesAttorneys                                     IssuingCardSpendingControlsBlockedCategory = "legal_services_attorneys"
	IssuingCardSpendingControlsBlockedCategoryLuggageAndLeatherGoodsStores                               IssuingCardSpendingControlsBlockedCategory = "luggage_and_leather_goods_stores"
	IssuingCardSpendingControlsBlockedCategoryLumberBuildingMaterialsStores                              IssuingCardSpendingControlsBlockedCategory = "lumber_building_materials_stores"
	IssuingCardSpendingControlsBlockedCategoryManualCashDisburse                                         IssuingCardSpendingControlsBlockedCategory = "manual_cash_disburse"
	IssuingCardSpendingControlsBlockedCategoryMarinasServiceAndSupplies                                  IssuingCardSpendingControlsBlockedCategory = "marinas_service_and_supplies"
	IssuingCardSpendingControlsBlockedCategoryMasonryStoneworkAndPlaster                                 IssuingCardSpendingControlsBlockedCategory = "masonry_stonework_and_plaster"
	IssuingCardSpendingControlsBlockedCategoryMassageParlors                                             IssuingCardSpendingControlsBlockedCategory = "massage_parlors"
	IssuingCardSpendingControlsBlockedCategoryMedicalAndDentalLabs                                       IssuingCardSpendingControlsBlockedCategory = "medical_and_dental_labs"
	IssuingCardSpendingControlsBlockedCategoryMedicalDentalOphthalmicAndHospitalEquipmentAndSupplies     IssuingCardSpendingControlsBlockedCategory = "medical_dental_ophthalmic_and_hospital_equipment_and_supplies"
	IssuingCardSpendingControlsBlockedCategoryMedicalServices                                            IssuingCardSpendingControlsBlockedCategory = "medical_services"
	IssuingCardSpendingControlsBlockedCategoryMembershipOrganizations                                    IssuingCardSpendingControlsBlockedCategory = "membership_organizations"
	IssuingCardSpendingControlsBlockedCategoryMensAndBoysClothingAndAccessoriesStores                    IssuingCardSpendingControlsBlockedCategory = "mens_and_boys_clothing_and_accessories_stores"
	IssuingCardSpendingControlsBlockedCategoryMensWomensClothingStores                                   IssuingCardSpendingControlsBlockedCategory = "mens_womens_clothing_stores"
	IssuingCardSpendingControlsBlockedCategoryMetalServiceCenters                                        IssuingCardSpendingControlsBlockedCategory = "metal_service_centers"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneous                                              IssuingCardSpendingControlsBlockedCategory = "miscellaneous"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousApparelAndAccessoryShops                      IssuingCardSpendingControlsBlockedCategory = "miscellaneous_apparel_and_accessory_shops"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousAutoDealers                                   IssuingCardSpendingControlsBlockedCategory = "miscellaneous_auto_dealers"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousBusinessServices                              IssuingCardSpendingControlsBlockedCategory = "miscellaneous_business_services"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousFoodStores                                    IssuingCardSpendingControlsBlockedCategory = "miscellaneous_food_stores"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousGeneralMerchandise                            IssuingCardSpendingControlsBlockedCategory = "miscellaneous_general_merchandise"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousGeneralServices                               IssuingCardSpendingControlsBlockedCategory = "miscellaneous_general_services"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousHomeFurnishingSpecialtyStores                 IssuingCardSpendingControlsBlockedCategory = "miscellaneous_home_furnishing_specialty_stores"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousPublishingAndPrinting                         IssuingCardSpendingControlsBlockedCategory = "miscellaneous_publishing_and_printing"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousRecreationServices                            IssuingCardSpendingControlsBlockedCategory = "miscellaneous_recreation_services"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousRepairShops                                   IssuingCardSpendingControlsBlockedCategory = "miscellaneous_repair_shops"
	IssuingCardSpendingControlsBlockedCategoryMiscellaneousSpecialtyRetail                               IssuingCardSpendingControlsBlockedCategory = "miscellaneous_specialty_retail"
	IssuingCardSpendingControlsBlockedCategoryMobileHomeDealers                                          IssuingCardSpendingControlsBlockedCategory = "mobile_home_dealers"
	IssuingCardSpendingControlsBlockedCategoryMotionPictureTheaters                                      IssuingCardSpendingControlsBlockedCategory = "motion_picture_theaters"
	IssuingCardSpendingControlsBlockedCategoryMotorFreightCarriersAndTrucking                            IssuingCardSpendingControlsBlockedCategory = "motor_freight_carriers_and_trucking"
	IssuingCardSpendingControlsBlockedCategoryMotorHomesDealers                                          IssuingCardSpendingControlsBlockedCategory = "motor_homes_dealers"
	IssuingCardSpendingControlsBlockedCategoryMotorVehicleSuppliesAndNewParts                            IssuingCardSpendingControlsBlockedCategory = "motor_vehicle_supplies_and_new_parts"
	IssuingCardSpendingControlsBlockedCategoryMotorcycleShopsAndDealers                                  IssuingCardSpendingControlsBlockedCategory = "motorcycle_shops_and_dealers"
	IssuingCardSpendingControlsBlockedCategoryMotorcycleShopsDealers                                     IssuingCardSpendingControlsBlockedCategory = "motorcycle_shops_dealers"
	IssuingCardSpendingControlsBlockedCategoryMusicStoresMusicalInstrumentsPianosAndSheetMusic           IssuingCardSpendingControlsBlockedCategory = "music_stores_musical_instruments_pianos_and_sheet_music"
	IssuingCardSpendingControlsBlockedCategoryNewsDealersAndNewsstands                                   IssuingCardSpendingControlsBlockedCategory = "news_dealers_and_newsstands"
	IssuingCardSpendingControlsBlockedCategoryNonFiMoneyOrders                                           IssuingCardSpendingControlsBlockedCategory = "non_fi_money_orders"
	IssuingCardSpendingControlsBlockedCategoryNonFiStoredValueCardPurchaseLoad                           IssuingCardSpendingControlsBlockedCategory = "non_fi_stored_value_card_purchase_load"
	IssuingCardSpendingControlsBlockedCategoryNondurableGoods                                            IssuingCardSpendingControlsBlockedCategory = "nondurable_goods"
	IssuingCardSpendingControlsBlockedCategoryNurseriesLawnAndGardenSupplyStores                         IssuingCardSpendingControlsBlockedCategory = "nurseries_lawn_and_garden_supply_stores"
	IssuingCardSpendingControlsBlockedCategoryNursingPersonalCare                                        IssuingCardSpendingControlsBlockedCategory = "nursing_personal_care"
	IssuingCardSpendingControlsBlockedCategoryOfficeAndCommercialFurniture                               IssuingCardSpendingControlsBlockedCategory = "office_and_commercial_furniture"
	IssuingCardSpendingControlsBlockedCategoryOpticiansEyeglasses                                        IssuingCardSpendingControlsBlockedCategory = "opticians_eyeglasses"
	IssuingCardSpendingControlsBlockedCategoryOptometristsOphthalmologist                                IssuingCardSpendingControlsBlockedCategory = "optometrists_ophthalmologist"
	IssuingCardSpendingControlsBlockedCategoryOrthopedicGoodsProstheticDevices                           IssuingCardSpendingControlsBlockedCategory = "orthopedic_goods_prosthetic_devices"
	IssuingCardSpendingControlsBlockedCategoryOsteopaths                                                 IssuingCardSpendingControlsBlockedCategory = "osteopaths"
	IssuingCardSpendingControlsBlockedCategoryPackageStoresBeerWineAndLiquor                             IssuingCardSpendingControlsBlockedCategory = "package_stores_beer_wine_and_liquor"
	IssuingCardSpendingControlsBlockedCategoryPaintsVarnishesAndSupplies                                 IssuingCardSpendingControlsBlockedCategory = "paints_varnishes_and_supplies"
	IssuingCardSpendingControlsBlockedCategoryParkingLotsGarages                                         IssuingCardSpendingControlsBlockedCategory = "parking_lots_garages"
	IssuingCardSpendingControlsBlockedCategoryPassengerRailways                                          IssuingCardSpendingControlsBlockedCategory = "passenger_railways"
	IssuingCardSpendingControlsBlockedCategoryPawnShops                                                  IssuingCardSpendingControlsBlockedCategory = "pawn_shops"
	IssuingCardSpendingControlsBlockedCategoryPetShopsPetFoodAndSupplies                                 IssuingCardSpendingControlsBlockedCategory = "pet_shops_pet_food_and_supplies"
	IssuingCardSpendingControlsBlockedCategoryPetroleumAndPetroleumProducts                              IssuingCardSpendingControlsBlockedCategory = "petroleum_and_petroleum_products"
	IssuingCardSpendingControlsBlockedCategoryPhotoDeveloping                                            IssuingCardSpendingControlsBlockedCategory = "photo_developing"
	IssuingCardSpendingControlsBlockedCategoryPhotographicPhotocopyMicrofilmEquipmentAndSupplies         IssuingCardSpendingControlsBlockedCategory = "photographic_photocopy_microfilm_equipment_and_supplies"
	IssuingCardSpendingControlsBlockedCategoryPhotographicStudios                                        IssuingCardSpendingControlsBlockedCategory = "photographic_studios"
	IssuingCardSpendingControlsBlockedCategoryPictureVideoProduction                                     IssuingCardSpendingControlsBlockedCategory = "picture_video_production"
	IssuingCardSpendingControlsBlockedCategoryPieceGoodsNotionsAndOtherDryGoods                          IssuingCardSpendingControlsBlockedCategory = "piece_goods_notions_and_other_dry_goods"
	IssuingCardSpendingControlsBlockedCategoryPlumbingHeatingEquipmentAndSupplies                        IssuingCardSpendingControlsBlockedCategory = "plumbing_heating_equipment_and_supplies"
	IssuingCardSpendingControlsBlockedCategoryPoliticalOrganizations                                     IssuingCardSpendingControlsBlockedCategory = "political_organizations"
	IssuingCardSpendingControlsBlockedCategoryPostalServicesGovernmentOnly                               IssuingCardSpendingControlsBlockedCategory = "postal_services_government_only"
	IssuingCardSpendingControlsBlockedCategoryPreciousStonesAndMetalsWatchesAndJewelry                   IssuingCardSpendingControlsBlockedCategory = "precious_stones_and_metals_watches_and_jewelry"
	IssuingCardSpendingControlsBlockedCategoryProfessionalServices                                       IssuingCardSpendingControlsBlockedCategory = "professional_services"
	IssuingCardSpendingControlsBlockedCategoryPublicWarehousingAndStorage                                IssuingCardSpendingControlsBlockedCategory = "public_warehousing_and_storage"
	IssuingCardSpendingControlsBlockedCategoryQuickCopyReproAndBlueprint                                 IssuingCardSpendingControlsBlockedCategory = "quick_copy_repro_and_blueprint"
	IssuingCardSpendingControlsBlockedCategoryRailroads                                                  IssuingCardSpendingControlsBlockedCategory = "railroads"
	IssuingCardSpendingControlsBlockedCategoryRealEstateAgentsAndManagersRentals                         IssuingCardSpendingControlsBlockedCategory = "real_estate_agents_and_managers_rentals"
	IssuingCardSpendingControlsBlockedCategoryRecordStores                                               IssuingCardSpendingControlsBlockedCategory = "record_stores"
	IssuingCardSpendingControlsBlockedCategoryRecreationalVehicleRentals                                 IssuingCardSpendingControlsBlockedCategory = "recreational_vehicle_rentals"
	IssuingCardSpendingControlsBlockedCategoryReligiousGoodsStores                                       IssuingCardSpendingControlsBlockedCategory = "religious_goods_stores"
	IssuingCardSpendingControlsBlockedCategoryReligiousOrganizations                                     IssuingCardSpendingControlsBlockedCategory = "religious_organizations"
	IssuingCardSpendingControlsBlockedCategoryRoofingSidingSheetMetal                                    IssuingCardSpendingControlsBlockedCategory = "roofing_siding_sheet_metal"
	IssuingCardSpendingControlsBlockedCategorySecretarialSupportServices                                 IssuingCardSpendingControlsBlockedCategory = "secretarial_support_services"
	IssuingCardSpendingControlsBlockedCategorySecurityBrokersDealers                                     IssuingCardSpendingControlsBlockedCategory = "security_brokers_dealers"
	IssuingCardSpendingControlsBlockedCategoryServiceStations                                            IssuingCardSpendingControlsBlockedCategory = "service_stations"
	IssuingCardSpendingControlsBlockedCategorySewingNeedleworkFabricAndPieceGoodsStores                  IssuingCardSpendingControlsBlockedCategory = "sewing_needlework_fabric_and_piece_goods_stores"
	IssuingCardSpendingControlsBlockedCategoryShoeRepairHatCleaning                                      IssuingCardSpendingControlsBlockedCategory = "shoe_repair_hat_cleaning"
	IssuingCardSpendingControlsBlockedCategoryShoeStores                                                 IssuingCardSpendingControlsBlockedCategory = "shoe_stores"
	IssuingCardSpendingControlsBlockedCategorySmallApplianceRepair                                       IssuingCardSpendingControlsBlockedCategory = "small_appliance_repair"
	IssuingCardSpendingControlsBlockedCategorySnowmobileDealers                                          IssuingCardSpendingControlsBlockedCategory = "snowmobile_dealers"
	IssuingCardSpendingControlsBlockedCategorySpecialTradeServices                                       IssuingCardSpendingControlsBlockedCategory = "special_trade_services"
	IssuingCardSpendingControlsBlockedCategorySpecialtyCleaning                                          IssuingCardSpendingControlsBlockedCategory = "specialty_cleaning"
	IssuingCardSpendingControlsBlockedCategorySportingGoodsStores                                        IssuingCardSpendingControlsBlockedCategory = "sporting_goods_stores"
	IssuingCardSpendingControlsBlockedCategorySportingRecreationCamps                                    IssuingCardSpendingControlsBlockedCategory = "sporting_recreation_camps"
	IssuingCardSpendingControlsBlockedCategorySportsAndRidingApparelStores                               IssuingCardSpendingControlsBlockedCategory = "sports_and_riding_apparel_stores"
	IssuingCardSpendingControlsBlockedCategorySportsClubsFields                                          IssuingCardSpendingControlsBlockedCategory = "sports_clubs_fields"
	IssuingCardSpendingControlsBlockedCategoryStampAndCoinStores                                         IssuingCardSpendingControlsBlockedCategory = "stamp_and_coin_stores"
	IssuingCardSpendingControlsBlockedCategoryStationaryOfficeSuppliesPrintingAndWritingPaper            IssuingCardSpendingControlsBlockedCategory = "stationary_office_supplies_printing_and_writing_paper"
	IssuingCardSpendingControlsBlockedCategoryStationeryStoresOfficeAndSchoolSupplyStores                IssuingCardSpendingControlsBlockedCategory = "stationery_stores_office_and_school_supply_stores"
	IssuingCardSpendingControlsBlockedCategorySwimmingPoolsSales                                         IssuingCardSpendingControlsBlockedCategory = "swimming_pools_sales"
	IssuingCardSpendingControlsBlockedCategoryTUiTravelGermany                                           IssuingCardSpendingControlsBlockedCategory = "t_ui_travel_germany"
	IssuingCardSpendingControlsBlockedCategoryTailorsAlterations                                         IssuingCardSpendingControlsBlockedCategory = "tailors_alterations"
	IssuingCardSpendingControlsBlockedCategoryTaxPaymentsGovernmentAgencies                              IssuingCardSpendingControlsBlockedCategory = "tax_payments_government_agencies"
	IssuingCardSpendingControlsBlockedCategoryTaxPreparationServices                                     IssuingCardSpendingControlsBlockedCategory = "tax_preparation_services"
	IssuingCardSpendingControlsBlockedCategoryTaxicabsLimousines                                         IssuingCardSpendingControlsBlockedCategory = "taxicabs_limousines"
	IssuingCardSpendingControlsBlockedCategoryTelecommunicationEquipmentAndTelephoneSales                IssuingCardSpendingControlsBlockedCategory = "telecommunication_equipment_and_telephone_sales"
	IssuingCardSpendingControlsBlockedCategoryTelecommunicationServices                                  IssuingCardSpendingControlsBlockedCategory = "telecommunication_services"
	IssuingCardSpendingControlsBlockedCategoryTelegraphServices                                          IssuingCardSpendingControlsBlockedCategory = "telegraph_services"
	IssuingCardSpendingControlsBlockedCategoryTentAndAwningShops                                         IssuingCardSpendingControlsBlockedCategory = "tent_and_awning_shops"
	IssuingCardSpendingControlsBlockedCategoryTestingLaboratories                                        IssuingCardSpendingControlsBlockedCategory = "testing_laboratories"
	IssuingCardSpendingControlsBlockedCategoryTheatricalTicketAgencies                                   IssuingCardSpendingControlsBlockedCategory = "theatrical_ticket_agencies"
	IssuingCardSpendingControlsBlockedCategoryTimeshares                                                 IssuingCardSpendingControlsBlockedCategory = "timeshares"
	IssuingCardSpendingControlsBlockedCategoryTireRetreadingAndRepair                                    IssuingCardSpendingControlsBlockedCategory = "tire_retreading_and_repair"
	IssuingCardSpendingControlsBlockedCategoryTollsBridgeFees                                            IssuingCardSpendingControlsBlockedCategory = "tolls_bridge_fees"
	IssuingCardSpendingControlsBlockedCategoryTouristAttractionsAndExhibits                              IssuingCardSpendingControlsBlockedCategory = "tourist_attractions_and_exhibits"
	IssuingCardSpendingControlsBlockedCategoryTowingServices                                             IssuingCardSpendingControlsBlockedCategory = "towing_services"
	IssuingCardSpendingControlsBlockedCategoryTrailerParksCampgrounds                                    IssuingCardSpendingControlsBlockedCategory = "trailer_parks_campgrounds"
	IssuingCardSpendingControlsBlockedCategoryTransportationServices                                     IssuingCardSpendingControlsBlockedCategory = "transportation_services"
	IssuingCardSpendingControlsBlockedCategoryTravelAgenciesTourOperators                                IssuingCardSpendingControlsBlockedCategory = "travel_agencies_tour_operators"
	IssuingCardSpendingControlsBlockedCategoryTruckStopIteration                                         IssuingCardSpendingControlsBlockedCategory = "truck_stop_iteration"
	IssuingCardSpendingControlsBlockedCategoryTruckUtilityTrailerRentals                                 IssuingCardSpendingControlsBlockedCategory = "truck_utility_trailer_rentals"
	IssuingCardSpendingControlsBlockedCategoryTypesettingPlateMakingAndRelatedServices                   IssuingCardSpendingControlsBlockedCategory = "typesetting_plate_making_and_related_services"
	IssuingCardSpendingControlsBlockedCategoryTypewriterStores                                           IssuingCardSpendingControlsBlockedCategory = "typewriter_stores"
	IssuingCardSpendingControlsBlockedCategoryUSFederalGovernmentAgenciesOrDepartments                   IssuingCardSpendingControlsBlockedCategory = "u_s_federal_government_agencies_or_departments"
	IssuingCardSpendingControlsBlockedCategoryUniformsCommercialClothing                                 IssuingCardSpendingControlsBlockedCategory = "uniforms_commercial_clothing"
	IssuingCardSpendingControlsBlockedCategoryUsedMerchandiseAndSecondhandStores                         IssuingCardSpendingControlsBlockedCategory = "used_merchandise_and_secondhand_stores"
	IssuingCardSpendingControlsBlockedCategoryUtilities                                                  IssuingCardSpendingControlsBlockedCategory = "utilities"
	IssuingCardSpendingControlsBlockedCategoryVarietyStores                                              IssuingCardSpendingControlsBlockedCategory = "variety_stores"
	IssuingCardSpendingControlsBlockedCategoryVeterinaryServices                                         IssuingCardSpendingControlsBlockedCategory = "veterinary_services"
	IssuingCardSpendingControlsBlockedCategoryVideoAmusementGameSupplies                                 IssuingCardSpendingControlsBlockedCategory = "video_amusement_game_supplies"
	IssuingCardSpendingControlsBlockedCategoryVideoGameArcades                                           IssuingCardSpendingControlsBlockedCategory = "video_game_arcades"
	IssuingCardSpendingControlsBlockedCategoryVideoTapeRentalStores                                      IssuingCardSpendingControlsBlockedCategory = "video_tape_rental_stores"
	IssuingCardSpendingControlsBlockedCategoryVocationalTradeSchools                                     IssuingCardSpendingControlsBlockedCategory = "vocational_trade_schools"
	IssuingCardSpendingControlsBlockedCategoryWatchJewelryRepair                                         IssuingCardSpendingControlsBlockedCategory = "watch_jewelry_repair"
	IssuingCardSpendingControlsBlockedCategoryWeldingRepair                                              IssuingCardSpendingControlsBlockedCategory = "welding_repair"
	IssuingCardSpendingControlsBlockedCategoryWholesaleClubs                                             IssuingCardSpendingControlsBlockedCategory = "wholesale_clubs"
	IssuingCardSpendingControlsBlockedCategoryWigAndToupeeStores                                         IssuingCardSpendingControlsBlockedCategory = "wig_and_toupee_stores"
	IssuingCardSpendingControlsBlockedCategoryWiresMoneyOrders                                           IssuingCardSpendingControlsBlockedCategory = "wires_money_orders"
	IssuingCardSpendingControlsBlockedCategoryWomensAccessoryAndSpecialtyShops                           IssuingCardSpendingControlsBlockedCategory = "womens_accessory_and_specialty_shops"
	IssuingCardSpendingControlsBlockedCategoryWomensReadyToWearStores                                    IssuingCardSpendingControlsBlockedCategory = "womens_ready_to_wear_stores"
	IssuingCardSpendingControlsBlockedCategoryWreckingAndSalvageYards                                    IssuingCardSpendingControlsBlockedCategory = "wrecking_and_salvage_yards"
)

// Array of strings containing [categories](https://stripe.com/docs/api#issuing_authorization_object-merchant_data-category) this limit applies to. Omitting this field will apply the limit to all categories.
type IssuingCardSpendingControlsSpendingLimitCategory string

// List of values that IssuingCardSpendingControlsSpendingLimitCategory can take
const (
	IssuingCardSpendingControlsSpendingLimitCategoryAcRefrigerationRepair                                      IssuingCardSpendingControlsSpendingLimitCategory = "ac_refrigeration_repair"
	IssuingCardSpendingControlsSpendingLimitCategoryAccountingBookkeepingServices                              IssuingCardSpendingControlsSpendingLimitCategory = "accounting_bookkeeping_services"
	IssuingCardSpendingControlsSpendingLimitCategoryAdvertisingServices                                        IssuingCardSpendingControlsSpendingLimitCategory = "advertising_services"
	IssuingCardSpendingControlsSpendingLimitCategoryAgriculturalCooperative                                    IssuingCardSpendingControlsSpendingLimitCategory = "agricultural_cooperative"
	IssuingCardSpendingControlsSpendingLimitCategoryAirlinesAirCarriers                                        IssuingCardSpendingControlsSpendingLimitCategory = "airlines_air_carriers"
	IssuingCardSpendingControlsSpendingLimitCategoryAirportsFlyingFields                                       IssuingCardSpendingControlsSpendingLimitCategory = "airports_flying_fields"
	IssuingCardSpendingControlsSpendingLimitCategoryAmbulanceServices                                          IssuingCardSpendingControlsSpendingLimitCategory = "ambulance_services"
	IssuingCardSpendingControlsSpendingLimitCategoryAmusementParksCarnivals                                    IssuingCardSpendingControlsSpendingLimitCategory = "amusement_parks_carnivals"
	IssuingCardSpendingControlsSpendingLimitCategoryAntiqueReproductions                                       IssuingCardSpendingControlsSpendingLimitCategory = "antique_reproductions"
	IssuingCardSpendingControlsSpendingLimitCategoryAntiqueShops                                               IssuingCardSpendingControlsSpendingLimitCategory = "antique_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryAquariums                                                  IssuingCardSpendingControlsSpendingLimitCategory = "aquariums"
	IssuingCardSpendingControlsSpendingLimitCategoryArchitecturalSurveyingServices                             IssuingCardSpendingControlsSpendingLimitCategory = "architectural_surveying_services"
	IssuingCardSpendingControlsSpendingLimitCategoryArtDealersAndGalleries                                     IssuingCardSpendingControlsSpendingLimitCategory = "art_dealers_and_galleries"
	IssuingCardSpendingControlsSpendingLimitCategoryArtistsSupplyAndCraftShops                                 IssuingCardSpendingControlsSpendingLimitCategory = "artists_supply_and_craft_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryAutoAndHomeSupplyStores                                    IssuingCardSpendingControlsSpendingLimitCategory = "auto_and_home_supply_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryAutoBodyRepairShops                                        IssuingCardSpendingControlsSpendingLimitCategory = "auto_body_repair_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryAutoPaintShops                                             IssuingCardSpendingControlsSpendingLimitCategory = "auto_paint_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryAutoServiceShops                                           IssuingCardSpendingControlsSpendingLimitCategory = "auto_service_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryAutomatedCashDisburse                                      IssuingCardSpendingControlsSpendingLimitCategory = "automated_cash_disburse"
	IssuingCardSpendingControlsSpendingLimitCategoryAutomatedFuelDispensers                                    IssuingCardSpendingControlsSpendingLimitCategory = "automated_fuel_dispensers"
	IssuingCardSpendingControlsSpendingLimitCategoryAutomobileAssociations                                     IssuingCardSpendingControlsSpendingLimitCategory = "automobile_associations"
	IssuingCardSpendingControlsSpendingLimitCategoryAutomotivePartsAndAccessoriesStores                        IssuingCardSpendingControlsSpendingLimitCategory = "automotive_parts_and_accessories_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryAutomotiveTireStores                                       IssuingCardSpendingControlsSpendingLimitCategory = "automotive_tire_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryBailAndBondPayments                                        IssuingCardSpendingControlsSpendingLimitCategory = "bail_and_bond_payments"
	IssuingCardSpendingControlsSpendingLimitCategoryBakeries                                                   IssuingCardSpendingControlsSpendingLimitCategory = "bakeries"
	IssuingCardSpendingControlsSpendingLimitCategoryBandsOrchestras                                            IssuingCardSpendingControlsSpendingLimitCategory = "bands_orchestras"
	IssuingCardSpendingControlsSpendingLimitCategoryBarberAndBeautyShops                                       IssuingCardSpendingControlsSpendingLimitCategory = "barber_and_beauty_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryBettingCasinoGambling                                      IssuingCardSpendingControlsSpendingLimitCategory = "betting_casino_gambling"
	IssuingCardSpendingControlsSpendingLimitCategoryBicycleShops                                               IssuingCardSpendingControlsSpendingLimitCategory = "bicycle_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryBilliardPoolEstablishments                                 IssuingCardSpendingControlsSpendingLimitCategory = "billiard_pool_establishments"
	IssuingCardSpendingControlsSpendingLimitCategoryBoatDealers                                                IssuingCardSpendingControlsSpendingLimitCategory = "boat_dealers"
	IssuingCardSpendingControlsSpendingLimitCategoryBoatRentalsAndLeases                                       IssuingCardSpendingControlsSpendingLimitCategory = "boat_rentals_and_leases"
	IssuingCardSpendingControlsSpendingLimitCategoryBookStores                                                 IssuingCardSpendingControlsSpendingLimitCategory = "book_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryBooksPeriodicalsAndNewspapers                              IssuingCardSpendingControlsSpendingLimitCategory = "books_periodicals_and_newspapers"
	IssuingCardSpendingControlsSpendingLimitCategoryBowlingAlleys                                              IssuingCardSpendingControlsSpendingLimitCategory = "bowling_alleys"
	IssuingCardSpendingControlsSpendingLimitCategoryBusLines                                                   IssuingCardSpendingControlsSpendingLimitCategory = "bus_lines"
	IssuingCardSpendingControlsSpendingLimitCategoryBusinessSecretarialSchools                                 IssuingCardSpendingControlsSpendingLimitCategory = "business_secretarial_schools"
	IssuingCardSpendingControlsSpendingLimitCategoryBuyingShoppingServices                                     IssuingCardSpendingControlsSpendingLimitCategory = "buying_shopping_services"
	IssuingCardSpendingControlsSpendingLimitCategoryCableSatelliteAndOtherPayTelevisionAndRadio                IssuingCardSpendingControlsSpendingLimitCategory = "cable_satellite_and_other_pay_television_and_radio"
	IssuingCardSpendingControlsSpendingLimitCategoryCameraAndPhotographicSupplyStores                          IssuingCardSpendingControlsSpendingLimitCategory = "camera_and_photographic_supply_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryCandyNutAndConfectioneryStores                             IssuingCardSpendingControlsSpendingLimitCategory = "candy_nut_and_confectionery_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryCarAndTruckDealersNewUsed                                  IssuingCardSpendingControlsSpendingLimitCategory = "car_and_truck_dealers_new_used"
	IssuingCardSpendingControlsSpendingLimitCategoryCarAndTruckDealersUsedOnly                                 IssuingCardSpendingControlsSpendingLimitCategory = "car_and_truck_dealers_used_only"
	IssuingCardSpendingControlsSpendingLimitCategoryCarRentalAgencies                                          IssuingCardSpendingControlsSpendingLimitCategory = "car_rental_agencies"
	IssuingCardSpendingControlsSpendingLimitCategoryCarWashes                                                  IssuingCardSpendingControlsSpendingLimitCategory = "car_washes"
	IssuingCardSpendingControlsSpendingLimitCategoryCarpentryServices                                          IssuingCardSpendingControlsSpendingLimitCategory = "carpentry_services"
	IssuingCardSpendingControlsSpendingLimitCategoryCarpetUpholsteryCleaning                                   IssuingCardSpendingControlsSpendingLimitCategory = "carpet_upholstery_cleaning"
	IssuingCardSpendingControlsSpendingLimitCategoryCaterers                                                   IssuingCardSpendingControlsSpendingLimitCategory = "caterers"
	IssuingCardSpendingControlsSpendingLimitCategoryCharitableAndSocialServiceOrganizationsFundraising         IssuingCardSpendingControlsSpendingLimitCategory = "charitable_and_social_service_organizations_fundraising"
	IssuingCardSpendingControlsSpendingLimitCategoryChemicalsAndAlliedProducts                                 IssuingCardSpendingControlsSpendingLimitCategory = "chemicals_and_allied_products"
	IssuingCardSpendingControlsSpendingLimitCategoryChildCareServices                                          IssuingCardSpendingControlsSpendingLimitCategory = "child_care_services"
	IssuingCardSpendingControlsSpendingLimitCategoryChildrensAndInfantsWearStores                              IssuingCardSpendingControlsSpendingLimitCategory = "childrens_and_infants_wear_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryChiropodistsPodiatrists                                    IssuingCardSpendingControlsSpendingLimitCategory = "chiropodists_podiatrists"
	IssuingCardSpendingControlsSpendingLimitCategoryChiropractors                                              IssuingCardSpendingControlsSpendingLimitCategory = "chiropractors"
	IssuingCardSpendingControlsSpendingLimitCategoryCigarStoresAndStands                                       IssuingCardSpendingControlsSpendingLimitCategory = "cigar_stores_and_stands"
	IssuingCardSpendingControlsSpendingLimitCategoryCivicSocialFraternalAssociations                           IssuingCardSpendingControlsSpendingLimitCategory = "civic_social_fraternal_associations"
	IssuingCardSpendingControlsSpendingLimitCategoryCleaningAndMaintenance                                     IssuingCardSpendingControlsSpendingLimitCategory = "cleaning_and_maintenance"
	IssuingCardSpendingControlsSpendingLimitCategoryClothingRental                                             IssuingCardSpendingControlsSpendingLimitCategory = "clothing_rental"
	IssuingCardSpendingControlsSpendingLimitCategoryCollegesUniversities                                       IssuingCardSpendingControlsSpendingLimitCategory = "colleges_universities"
	IssuingCardSpendingControlsSpendingLimitCategoryCommercialEquipment                                        IssuingCardSpendingControlsSpendingLimitCategory = "commercial_equipment"
	IssuingCardSpendingControlsSpendingLimitCategoryCommercialFootwear                                         IssuingCardSpendingControlsSpendingLimitCategory = "commercial_footwear"
	IssuingCardSpendingControlsSpendingLimitCategoryCommercialPhotographyArtAndGraphics                        IssuingCardSpendingControlsSpendingLimitCategory = "commercial_photography_art_and_graphics"
	IssuingCardSpendingControlsSpendingLimitCategoryCommuterTransportAndFerries                                IssuingCardSpendingControlsSpendingLimitCategory = "commuter_transport_and_ferries"
	IssuingCardSpendingControlsSpendingLimitCategoryComputerNetworkServices                                    IssuingCardSpendingControlsSpendingLimitCategory = "computer_network_services"
	IssuingCardSpendingControlsSpendingLimitCategoryComputerProgramming                                        IssuingCardSpendingControlsSpendingLimitCategory = "computer_programming"
	IssuingCardSpendingControlsSpendingLimitCategoryComputerRepair                                             IssuingCardSpendingControlsSpendingLimitCategory = "computer_repair"
	IssuingCardSpendingControlsSpendingLimitCategoryComputerSoftwareStores                                     IssuingCardSpendingControlsSpendingLimitCategory = "computer_software_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryComputersPeripheralsAndSoftware                            IssuingCardSpendingControlsSpendingLimitCategory = "computers_peripherals_and_software"
	IssuingCardSpendingControlsSpendingLimitCategoryConcreteWorkServices                                       IssuingCardSpendingControlsSpendingLimitCategory = "concrete_work_services"
	IssuingCardSpendingControlsSpendingLimitCategoryConstructionMaterials                                      IssuingCardSpendingControlsSpendingLimitCategory = "construction_materials"
	IssuingCardSpendingControlsSpendingLimitCategoryConsultingPublicRelations                                  IssuingCardSpendingControlsSpendingLimitCategory = "consulting_public_relations"
	IssuingCardSpendingControlsSpendingLimitCategoryCorrespondenceSchools                                      IssuingCardSpendingControlsSpendingLimitCategory = "correspondence_schools"
	IssuingCardSpendingControlsSpendingLimitCategoryCosmeticStores                                             IssuingCardSpendingControlsSpendingLimitCategory = "cosmetic_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryCounselingServices                                         IssuingCardSpendingControlsSpendingLimitCategory = "counseling_services"
	IssuingCardSpendingControlsSpendingLimitCategoryCountryClubs                                               IssuingCardSpendingControlsSpendingLimitCategory = "country_clubs"
	IssuingCardSpendingControlsSpendingLimitCategoryCourierServices                                            IssuingCardSpendingControlsSpendingLimitCategory = "courier_services"
	IssuingCardSpendingControlsSpendingLimitCategoryCourtCosts                                                 IssuingCardSpendingControlsSpendingLimitCategory = "court_costs"
	IssuingCardSpendingControlsSpendingLimitCategoryCreditReportingAgencies                                    IssuingCardSpendingControlsSpendingLimitCategory = "credit_reporting_agencies"
	IssuingCardSpendingControlsSpendingLimitCategoryCruiseLines                                                IssuingCardSpendingControlsSpendingLimitCategory = "cruise_lines"
	IssuingCardSpendingControlsSpendingLimitCategoryDairyProductsStores                                        IssuingCardSpendingControlsSpendingLimitCategory = "dairy_products_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryDanceHallStudiosSchools                                    IssuingCardSpendingControlsSpendingLimitCategory = "dance_hall_studios_schools"
	IssuingCardSpendingControlsSpendingLimitCategoryDatingEscortServices                                       IssuingCardSpendingControlsSpendingLimitCategory = "dating_escort_services"
	IssuingCardSpendingControlsSpendingLimitCategoryDentistsOrthodontists                                      IssuingCardSpendingControlsSpendingLimitCategory = "dentists_orthodontists"
	IssuingCardSpendingControlsSpendingLimitCategoryDepartmentStores                                           IssuingCardSpendingControlsSpendingLimitCategory = "department_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryDetectiveAgencies                                          IssuingCardSpendingControlsSpendingLimitCategory = "detective_agencies"
	IssuingCardSpendingControlsSpendingLimitCategoryDigitalGoodsApplications                                   IssuingCardSpendingControlsSpendingLimitCategory = "digital_goods_applications"
	IssuingCardSpendingControlsSpendingLimitCategoryDigitalGoodsGames                                          IssuingCardSpendingControlsSpendingLimitCategory = "digital_goods_games"
	IssuingCardSpendingControlsSpendingLimitCategoryDigitalGoodsLargeVolume                                    IssuingCardSpendingControlsSpendingLimitCategory = "digital_goods_large_volume"
	IssuingCardSpendingControlsSpendingLimitCategoryDigitalGoodsMedia                                          IssuingCardSpendingControlsSpendingLimitCategory = "digital_goods_media"
	IssuingCardSpendingControlsSpendingLimitCategoryDirectMarketingCatalogMerchant                             IssuingCardSpendingControlsSpendingLimitCategory = "direct_marketing_catalog_merchant"
	IssuingCardSpendingControlsSpendingLimitCategoryDirectMarketingCombinationCatalogAndRetailMerchant         IssuingCardSpendingControlsSpendingLimitCategory = "direct_marketing_combination_catalog_and_retail_merchant"
	IssuingCardSpendingControlsSpendingLimitCategoryDirectMarketingInboundTelemarketing                        IssuingCardSpendingControlsSpendingLimitCategory = "direct_marketing_inbound_telemarketing"
	IssuingCardSpendingControlsSpendingLimitCategoryDirectMarketingInsuranceServices                           IssuingCardSpendingControlsSpendingLimitCategory = "direct_marketing_insurance_services"
	IssuingCardSpendingControlsSpendingLimitCategoryDirectMarketingOther                                       IssuingCardSpendingControlsSpendingLimitCategory = "direct_marketing_other"
	IssuingCardSpendingControlsSpendingLimitCategoryDirectMarketingOutboundTelemarketing                       IssuingCardSpendingControlsSpendingLimitCategory = "direct_marketing_outbound_telemarketing"
	IssuingCardSpendingControlsSpendingLimitCategoryDirectMarketingSubscription                                IssuingCardSpendingControlsSpendingLimitCategory = "direct_marketing_subscription"
	IssuingCardSpendingControlsSpendingLimitCategoryDirectMarketingTravel                                      IssuingCardSpendingControlsSpendingLimitCategory = "direct_marketing_travel"
	IssuingCardSpendingControlsSpendingLimitCategoryDiscountStores                                             IssuingCardSpendingControlsSpendingLimitCategory = "discount_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryDoctors                                                    IssuingCardSpendingControlsSpendingLimitCategory = "doctors"
	IssuingCardSpendingControlsSpendingLimitCategoryDoorToDoorSales                                            IssuingCardSpendingControlsSpendingLimitCategory = "door_to_door_sales"
	IssuingCardSpendingControlsSpendingLimitCategoryDraperyWindowCoveringAndUpholsteryStores                   IssuingCardSpendingControlsSpendingLimitCategory = "drapery_window_covering_and_upholstery_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryDrinkingPlaces                                             IssuingCardSpendingControlsSpendingLimitCategory = "drinking_places"
	IssuingCardSpendingControlsSpendingLimitCategoryDrugStoresAndPharmacies                                    IssuingCardSpendingControlsSpendingLimitCategory = "drug_stores_and_pharmacies"
	IssuingCardSpendingControlsSpendingLimitCategoryDrugsDrugProprietariesAndDruggistSundries                  IssuingCardSpendingControlsSpendingLimitCategory = "drugs_drug_proprietaries_and_druggist_sundries"
	IssuingCardSpendingControlsSpendingLimitCategoryDryCleaners                                                IssuingCardSpendingControlsSpendingLimitCategory = "dry_cleaners"
	IssuingCardSpendingControlsSpendingLimitCategoryDurableGoods                                               IssuingCardSpendingControlsSpendingLimitCategory = "durable_goods"
	IssuingCardSpendingControlsSpendingLimitCategoryDutyFreeStores                                             IssuingCardSpendingControlsSpendingLimitCategory = "duty_free_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryEatingPlacesRestaurants                                    IssuingCardSpendingControlsSpendingLimitCategory = "eating_places_restaurants"
	IssuingCardSpendingControlsSpendingLimitCategoryEducationalServices                                        IssuingCardSpendingControlsSpendingLimitCategory = "educational_services"
	IssuingCardSpendingControlsSpendingLimitCategoryElectricRazorStores                                        IssuingCardSpendingControlsSpendingLimitCategory = "electric_razor_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryElectricalPartsAndEquipment                                IssuingCardSpendingControlsSpendingLimitCategory = "electrical_parts_and_equipment"
	IssuingCardSpendingControlsSpendingLimitCategoryElectricalServices                                         IssuingCardSpendingControlsSpendingLimitCategory = "electrical_services"
	IssuingCardSpendingControlsSpendingLimitCategoryElectronicsRepairShops                                     IssuingCardSpendingControlsSpendingLimitCategory = "electronics_repair_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryElectronicsStores                                          IssuingCardSpendingControlsSpendingLimitCategory = "electronics_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryElementarySecondarySchools                                 IssuingCardSpendingControlsSpendingLimitCategory = "elementary_secondary_schools"
	IssuingCardSpendingControlsSpendingLimitCategoryEmploymentTempAgencies                                     IssuingCardSpendingControlsSpendingLimitCategory = "employment_temp_agencies"
	IssuingCardSpendingControlsSpendingLimitCategoryEquipmentRental                                            IssuingCardSpendingControlsSpendingLimitCategory = "equipment_rental"
	IssuingCardSpendingControlsSpendingLimitCategoryExterminatingServices                                      IssuingCardSpendingControlsSpendingLimitCategory = "exterminating_services"
	IssuingCardSpendingControlsSpendingLimitCategoryFamilyClothingStores                                       IssuingCardSpendingControlsSpendingLimitCategory = "family_clothing_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryFastFoodRestaurants                                        IssuingCardSpendingControlsSpendingLimitCategory = "fast_food_restaurants"
	IssuingCardSpendingControlsSpendingLimitCategoryFinancialInstitutions                                      IssuingCardSpendingControlsSpendingLimitCategory = "financial_institutions"
	IssuingCardSpendingControlsSpendingLimitCategoryFinesGovernmentAdministrativeEntities                      IssuingCardSpendingControlsSpendingLimitCategory = "fines_government_administrative_entities"
	IssuingCardSpendingControlsSpendingLimitCategoryFireplaceFireplaceScreensAndAccessoriesStores              IssuingCardSpendingControlsSpendingLimitCategory = "fireplace_fireplace_screens_and_accessories_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryFloorCoveringStores                                        IssuingCardSpendingControlsSpendingLimitCategory = "floor_covering_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryFlorists                                                   IssuingCardSpendingControlsSpendingLimitCategory = "florists"
	IssuingCardSpendingControlsSpendingLimitCategoryFloristsSuppliesNurseryStockAndFlowers                     IssuingCardSpendingControlsSpendingLimitCategory = "florists_supplies_nursery_stock_and_flowers"
	IssuingCardSpendingControlsSpendingLimitCategoryFreezerAndLockerMeatProvisioners                           IssuingCardSpendingControlsSpendingLimitCategory = "freezer_and_locker_meat_provisioners"
	IssuingCardSpendingControlsSpendingLimitCategoryFuelDealersNonAutomotive                                   IssuingCardSpendingControlsSpendingLimitCategory = "fuel_dealers_non_automotive"
	IssuingCardSpendingControlsSpendingLimitCategoryFuneralServicesCrematories                                 IssuingCardSpendingControlsSpendingLimitCategory = "funeral_services_crematories"
	IssuingCardSpendingControlsSpendingLimitCategoryFurnitureHomeFurnishingsAndEquipmentStoresExceptAppliances IssuingCardSpendingControlsSpendingLimitCategory = "furniture_home_furnishings_and_equipment_stores_except_appliances"
	IssuingCardSpendingControlsSpendingLimitCategoryFurnitureRepairRefinishing                                 IssuingCardSpendingControlsSpendingLimitCategory = "furniture_repair_refinishing"
	IssuingCardSpendingControlsSpendingLimitCategoryFurriersAndFurShops                                        IssuingCardSpendingControlsSpendingLimitCategory = "furriers_and_fur_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryGeneralServices                                            IssuingCardSpendingControlsSpendingLimitCategory = "general_services"
	IssuingCardSpendingControlsSpendingLimitCategoryGiftCardNoveltyAndSouvenirShops                            IssuingCardSpendingControlsSpendingLimitCategory = "gift_card_novelty_and_souvenir_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryGlassPaintAndWallpaperStores                               IssuingCardSpendingControlsSpendingLimitCategory = "glass_paint_and_wallpaper_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryGlasswareCrystalStores                                     IssuingCardSpendingControlsSpendingLimitCategory = "glassware_crystal_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryGolfCoursesPublic                                          IssuingCardSpendingControlsSpendingLimitCategory = "golf_courses_public"
	IssuingCardSpendingControlsSpendingLimitCategoryGovernmentServices                                         IssuingCardSpendingControlsSpendingLimitCategory = "government_services"
	IssuingCardSpendingControlsSpendingLimitCategoryGroceryStoresSupermarkets                                  IssuingCardSpendingControlsSpendingLimitCategory = "grocery_stores_supermarkets"
	IssuingCardSpendingControlsSpendingLimitCategoryHardwareEquipmentAndSupplies                               IssuingCardSpendingControlsSpendingLimitCategory = "hardware_equipment_and_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryHardwareStores                                             IssuingCardSpendingControlsSpendingLimitCategory = "hardware_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryHealthAndBeautySpas                                        IssuingCardSpendingControlsSpendingLimitCategory = "health_and_beauty_spas"
	IssuingCardSpendingControlsSpendingLimitCategoryHearingAidsSalesAndSupplies                                IssuingCardSpendingControlsSpendingLimitCategory = "hearing_aids_sales_and_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryHeatingPlumbingAC                                          IssuingCardSpendingControlsSpendingLimitCategory = "heating_plumbing_a_c"
	IssuingCardSpendingControlsSpendingLimitCategoryHobbyToyAndGameShops                                       IssuingCardSpendingControlsSpendingLimitCategory = "hobby_toy_and_game_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryHomeSupplyWarehouseStores                                  IssuingCardSpendingControlsSpendingLimitCategory = "home_supply_warehouse_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryHospitals                                                  IssuingCardSpendingControlsSpendingLimitCategory = "hospitals"
	IssuingCardSpendingControlsSpendingLimitCategoryHotelsMotelsAndResorts                                     IssuingCardSpendingControlsSpendingLimitCategory = "hotels_motels_and_resorts"
	IssuingCardSpendingControlsSpendingLimitCategoryHouseholdApplianceStores                                   IssuingCardSpendingControlsSpendingLimitCategory = "household_appliance_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryIndustrialSupplies                                         IssuingCardSpendingControlsSpendingLimitCategory = "industrial_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryInformationRetrievalServices                               IssuingCardSpendingControlsSpendingLimitCategory = "information_retrieval_services"
	IssuingCardSpendingControlsSpendingLimitCategoryInsuranceDefault                                           IssuingCardSpendingControlsSpendingLimitCategory = "insurance_default"
	IssuingCardSpendingControlsSpendingLimitCategoryInsuranceUnderwritingPremiums                              IssuingCardSpendingControlsSpendingLimitCategory = "insurance_underwriting_premiums"
	IssuingCardSpendingControlsSpendingLimitCategoryIntraCompanyPurchases                                      IssuingCardSpendingControlsSpendingLimitCategory = "intra_company_purchases"
	IssuingCardSpendingControlsSpendingLimitCategoryJewelryStoresWatchesClocksAndSilverwareStores              IssuingCardSpendingControlsSpendingLimitCategory = "jewelry_stores_watches_clocks_and_silverware_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryLandscapingServices                                        IssuingCardSpendingControlsSpendingLimitCategory = "landscaping_services"
	IssuingCardSpendingControlsSpendingLimitCategoryLaundries                                                  IssuingCardSpendingControlsSpendingLimitCategory = "laundries"
	IssuingCardSpendingControlsSpendingLimitCategoryLaundryCleaningServices                                    IssuingCardSpendingControlsSpendingLimitCategory = "laundry_cleaning_services"
	IssuingCardSpendingControlsSpendingLimitCategoryLegalServicesAttorneys                                     IssuingCardSpendingControlsSpendingLimitCategory = "legal_services_attorneys"
	IssuingCardSpendingControlsSpendingLimitCategoryLuggageAndLeatherGoodsStores                               IssuingCardSpendingControlsSpendingLimitCategory = "luggage_and_leather_goods_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryLumberBuildingMaterialsStores                              IssuingCardSpendingControlsSpendingLimitCategory = "lumber_building_materials_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryManualCashDisburse                                         IssuingCardSpendingControlsSpendingLimitCategory = "manual_cash_disburse"
	IssuingCardSpendingControlsSpendingLimitCategoryMarinasServiceAndSupplies                                  IssuingCardSpendingControlsSpendingLimitCategory = "marinas_service_and_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryMasonryStoneworkAndPlaster                                 IssuingCardSpendingControlsSpendingLimitCategory = "masonry_stonework_and_plaster"
	IssuingCardSpendingControlsSpendingLimitCategoryMassageParlors                                             IssuingCardSpendingControlsSpendingLimitCategory = "massage_parlors"
	IssuingCardSpendingControlsSpendingLimitCategoryMedicalAndDentalLabs                                       IssuingCardSpendingControlsSpendingLimitCategory = "medical_and_dental_labs"
	IssuingCardSpendingControlsSpendingLimitCategoryMedicalDentalOphthalmicAndHospitalEquipmentAndSupplies     IssuingCardSpendingControlsSpendingLimitCategory = "medical_dental_ophthalmic_and_hospital_equipment_and_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryMedicalServices                                            IssuingCardSpendingControlsSpendingLimitCategory = "medical_services"
	IssuingCardSpendingControlsSpendingLimitCategoryMembershipOrganizations                                    IssuingCardSpendingControlsSpendingLimitCategory = "membership_organizations"
	IssuingCardSpendingControlsSpendingLimitCategoryMensAndBoysClothingAndAccessoriesStores                    IssuingCardSpendingControlsSpendingLimitCategory = "mens_and_boys_clothing_and_accessories_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryMensWomensClothingStores                                   IssuingCardSpendingControlsSpendingLimitCategory = "mens_womens_clothing_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryMetalServiceCenters                                        IssuingCardSpendingControlsSpendingLimitCategory = "metal_service_centers"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneous                                              IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousApparelAndAccessoryShops                      IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_apparel_and_accessory_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousAutoDealers                                   IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_auto_dealers"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousBusinessServices                              IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_business_services"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousFoodStores                                    IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_food_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousGeneralMerchandise                            IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_general_merchandise"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousGeneralServices                               IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_general_services"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousHomeFurnishingSpecialtyStores                 IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_home_furnishing_specialty_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousPublishingAndPrinting                         IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_publishing_and_printing"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousRecreationServices                            IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_recreation_services"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousRepairShops                                   IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_repair_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryMiscellaneousSpecialtyRetail                               IssuingCardSpendingControlsSpendingLimitCategory = "miscellaneous_specialty_retail"
	IssuingCardSpendingControlsSpendingLimitCategoryMobileHomeDealers                                          IssuingCardSpendingControlsSpendingLimitCategory = "mobile_home_dealers"
	IssuingCardSpendingControlsSpendingLimitCategoryMotionPictureTheaters                                      IssuingCardSpendingControlsSpendingLimitCategory = "motion_picture_theaters"
	IssuingCardSpendingControlsSpendingLimitCategoryMotorFreightCarriersAndTrucking                            IssuingCardSpendingControlsSpendingLimitCategory = "motor_freight_carriers_and_trucking"
	IssuingCardSpendingControlsSpendingLimitCategoryMotorHomesDealers                                          IssuingCardSpendingControlsSpendingLimitCategory = "motor_homes_dealers"
	IssuingCardSpendingControlsSpendingLimitCategoryMotorVehicleSuppliesAndNewParts                            IssuingCardSpendingControlsSpendingLimitCategory = "motor_vehicle_supplies_and_new_parts"
	IssuingCardSpendingControlsSpendingLimitCategoryMotorcycleShopsAndDealers                                  IssuingCardSpendingControlsSpendingLimitCategory = "motorcycle_shops_and_dealers"
	IssuingCardSpendingControlsSpendingLimitCategoryMotorcycleShopsDealers                                     IssuingCardSpendingControlsSpendingLimitCategory = "motorcycle_shops_dealers"
	IssuingCardSpendingControlsSpendingLimitCategoryMusicStoresMusicalInstrumentsPianosAndSheetMusic           IssuingCardSpendingControlsSpendingLimitCategory = "music_stores_musical_instruments_pianos_and_sheet_music"
	IssuingCardSpendingControlsSpendingLimitCategoryNewsDealersAndNewsstands                                   IssuingCardSpendingControlsSpendingLimitCategory = "news_dealers_and_newsstands"
	IssuingCardSpendingControlsSpendingLimitCategoryNonFiMoneyOrders                                           IssuingCardSpendingControlsSpendingLimitCategory = "non_fi_money_orders"
	IssuingCardSpendingControlsSpendingLimitCategoryNonFiStoredValueCardPurchaseLoad                           IssuingCardSpendingControlsSpendingLimitCategory = "non_fi_stored_value_card_purchase_load"
	IssuingCardSpendingControlsSpendingLimitCategoryNondurableGoods                                            IssuingCardSpendingControlsSpendingLimitCategory = "nondurable_goods"
	IssuingCardSpendingControlsSpendingLimitCategoryNurseriesLawnAndGardenSupplyStores                         IssuingCardSpendingControlsSpendingLimitCategory = "nurseries_lawn_and_garden_supply_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryNursingPersonalCare                                        IssuingCardSpendingControlsSpendingLimitCategory = "nursing_personal_care"
	IssuingCardSpendingControlsSpendingLimitCategoryOfficeAndCommercialFurniture                               IssuingCardSpendingControlsSpendingLimitCategory = "office_and_commercial_furniture"
	IssuingCardSpendingControlsSpendingLimitCategoryOpticiansEyeglasses                                        IssuingCardSpendingControlsSpendingLimitCategory = "opticians_eyeglasses"
	IssuingCardSpendingControlsSpendingLimitCategoryOptometristsOphthalmologist                                IssuingCardSpendingControlsSpendingLimitCategory = "optometrists_ophthalmologist"
	IssuingCardSpendingControlsSpendingLimitCategoryOrthopedicGoodsProstheticDevices                           IssuingCardSpendingControlsSpendingLimitCategory = "orthopedic_goods_prosthetic_devices"
	IssuingCardSpendingControlsSpendingLimitCategoryOsteopaths                                                 IssuingCardSpendingControlsSpendingLimitCategory = "osteopaths"
	IssuingCardSpendingControlsSpendingLimitCategoryPackageStoresBeerWineAndLiquor                             IssuingCardSpendingControlsSpendingLimitCategory = "package_stores_beer_wine_and_liquor"
	IssuingCardSpendingControlsSpendingLimitCategoryPaintsVarnishesAndSupplies                                 IssuingCardSpendingControlsSpendingLimitCategory = "paints_varnishes_and_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryParkingLotsGarages                                         IssuingCardSpendingControlsSpendingLimitCategory = "parking_lots_garages"
	IssuingCardSpendingControlsSpendingLimitCategoryPassengerRailways                                          IssuingCardSpendingControlsSpendingLimitCategory = "passenger_railways"
	IssuingCardSpendingControlsSpendingLimitCategoryPawnShops                                                  IssuingCardSpendingControlsSpendingLimitCategory = "pawn_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryPetShopsPetFoodAndSupplies                                 IssuingCardSpendingControlsSpendingLimitCategory = "pet_shops_pet_food_and_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryPetroleumAndPetroleumProducts                              IssuingCardSpendingControlsSpendingLimitCategory = "petroleum_and_petroleum_products"
	IssuingCardSpendingControlsSpendingLimitCategoryPhotoDeveloping                                            IssuingCardSpendingControlsSpendingLimitCategory = "photo_developing"
	IssuingCardSpendingControlsSpendingLimitCategoryPhotographicPhotocopyMicrofilmEquipmentAndSupplies         IssuingCardSpendingControlsSpendingLimitCategory = "photographic_photocopy_microfilm_equipment_and_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryPhotographicStudios                                        IssuingCardSpendingControlsSpendingLimitCategory = "photographic_studios"
	IssuingCardSpendingControlsSpendingLimitCategoryPictureVideoProduction                                     IssuingCardSpendingControlsSpendingLimitCategory = "picture_video_production"
	IssuingCardSpendingControlsSpendingLimitCategoryPieceGoodsNotionsAndOtherDryGoods                          IssuingCardSpendingControlsSpendingLimitCategory = "piece_goods_notions_and_other_dry_goods"
	IssuingCardSpendingControlsSpendingLimitCategoryPlumbingHeatingEquipmentAndSupplies                        IssuingCardSpendingControlsSpendingLimitCategory = "plumbing_heating_equipment_and_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryPoliticalOrganizations                                     IssuingCardSpendingControlsSpendingLimitCategory = "political_organizations"
	IssuingCardSpendingControlsSpendingLimitCategoryPostalServicesGovernmentOnly                               IssuingCardSpendingControlsSpendingLimitCategory = "postal_services_government_only"
	IssuingCardSpendingControlsSpendingLimitCategoryPreciousStonesAndMetalsWatchesAndJewelry                   IssuingCardSpendingControlsSpendingLimitCategory = "precious_stones_and_metals_watches_and_jewelry"
	IssuingCardSpendingControlsSpendingLimitCategoryProfessionalServices                                       IssuingCardSpendingControlsSpendingLimitCategory = "professional_services"
	IssuingCardSpendingControlsSpendingLimitCategoryPublicWarehousingAndStorage                                IssuingCardSpendingControlsSpendingLimitCategory = "public_warehousing_and_storage"
	IssuingCardSpendingControlsSpendingLimitCategoryQuickCopyReproAndBlueprint                                 IssuingCardSpendingControlsSpendingLimitCategory = "quick_copy_repro_and_blueprint"
	IssuingCardSpendingControlsSpendingLimitCategoryRailroads                                                  IssuingCardSpendingControlsSpendingLimitCategory = "railroads"
	IssuingCardSpendingControlsSpendingLimitCategoryRealEstateAgentsAndManagersRentals                         IssuingCardSpendingControlsSpendingLimitCategory = "real_estate_agents_and_managers_rentals"
	IssuingCardSpendingControlsSpendingLimitCategoryRecordStores                                               IssuingCardSpendingControlsSpendingLimitCategory = "record_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryRecreationalVehicleRentals                                 IssuingCardSpendingControlsSpendingLimitCategory = "recreational_vehicle_rentals"
	IssuingCardSpendingControlsSpendingLimitCategoryReligiousGoodsStores                                       IssuingCardSpendingControlsSpendingLimitCategory = "religious_goods_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryReligiousOrganizations                                     IssuingCardSpendingControlsSpendingLimitCategory = "religious_organizations"
	IssuingCardSpendingControlsSpendingLimitCategoryRoofingSidingSheetMetal                                    IssuingCardSpendingControlsSpendingLimitCategory = "roofing_siding_sheet_metal"
	IssuingCardSpendingControlsSpendingLimitCategorySecretarialSupportServices                                 IssuingCardSpendingControlsSpendingLimitCategory = "secretarial_support_services"
	IssuingCardSpendingControlsSpendingLimitCategorySecurityBrokersDealers                                     IssuingCardSpendingControlsSpendingLimitCategory = "security_brokers_dealers"
	IssuingCardSpendingControlsSpendingLimitCategoryServiceStations                                            IssuingCardSpendingControlsSpendingLimitCategory = "service_stations"
	IssuingCardSpendingControlsSpendingLimitCategorySewingNeedleworkFabricAndPieceGoodsStores                  IssuingCardSpendingControlsSpendingLimitCategory = "sewing_needlework_fabric_and_piece_goods_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryShoeRepairHatCleaning                                      IssuingCardSpendingControlsSpendingLimitCategory = "shoe_repair_hat_cleaning"
	IssuingCardSpendingControlsSpendingLimitCategoryShoeStores                                                 IssuingCardSpendingControlsSpendingLimitCategory = "shoe_stores"
	IssuingCardSpendingControlsSpendingLimitCategorySmallApplianceRepair                                       IssuingCardSpendingControlsSpendingLimitCategory = "small_appliance_repair"
	IssuingCardSpendingControlsSpendingLimitCategorySnowmobileDealers                                          IssuingCardSpendingControlsSpendingLimitCategory = "snowmobile_dealers"
	IssuingCardSpendingControlsSpendingLimitCategorySpecialTradeServices                                       IssuingCardSpendingControlsSpendingLimitCategory = "special_trade_services"
	IssuingCardSpendingControlsSpendingLimitCategorySpecialtyCleaning                                          IssuingCardSpendingControlsSpendingLimitCategory = "specialty_cleaning"
	IssuingCardSpendingControlsSpendingLimitCategorySportingGoodsStores                                        IssuingCardSpendingControlsSpendingLimitCategory = "sporting_goods_stores"
	IssuingCardSpendingControlsSpendingLimitCategorySportingRecreationCamps                                    IssuingCardSpendingControlsSpendingLimitCategory = "sporting_recreation_camps"
	IssuingCardSpendingControlsSpendingLimitCategorySportsAndRidingApparelStores                               IssuingCardSpendingControlsSpendingLimitCategory = "sports_and_riding_apparel_stores"
	IssuingCardSpendingControlsSpendingLimitCategorySportsClubsFields                                          IssuingCardSpendingControlsSpendingLimitCategory = "sports_clubs_fields"
	IssuingCardSpendingControlsSpendingLimitCategoryStampAndCoinStores                                         IssuingCardSpendingControlsSpendingLimitCategory = "stamp_and_coin_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryStationaryOfficeSuppliesPrintingAndWritingPaper            IssuingCardSpendingControlsSpendingLimitCategory = "stationary_office_supplies_printing_and_writing_paper"
	IssuingCardSpendingControlsSpendingLimitCategoryStationeryStoresOfficeAndSchoolSupplyStores                IssuingCardSpendingControlsSpendingLimitCategory = "stationery_stores_office_and_school_supply_stores"
	IssuingCardSpendingControlsSpendingLimitCategorySwimmingPoolsSales                                         IssuingCardSpendingControlsSpendingLimitCategory = "swimming_pools_sales"
	IssuingCardSpendingControlsSpendingLimitCategoryTUiTravelGermany                                           IssuingCardSpendingControlsSpendingLimitCategory = "t_ui_travel_germany"
	IssuingCardSpendingControlsSpendingLimitCategoryTailorsAlterations                                         IssuingCardSpendingControlsSpendingLimitCategory = "tailors_alterations"
	IssuingCardSpendingControlsSpendingLimitCategoryTaxPaymentsGovernmentAgencies                              IssuingCardSpendingControlsSpendingLimitCategory = "tax_payments_government_agencies"
	IssuingCardSpendingControlsSpendingLimitCategoryTaxPreparationServices                                     IssuingCardSpendingControlsSpendingLimitCategory = "tax_preparation_services"
	IssuingCardSpendingControlsSpendingLimitCategoryTaxicabsLimousines                                         IssuingCardSpendingControlsSpendingLimitCategory = "taxicabs_limousines"
	IssuingCardSpendingControlsSpendingLimitCategoryTelecommunicationEquipmentAndTelephoneSales                IssuingCardSpendingControlsSpendingLimitCategory = "telecommunication_equipment_and_telephone_sales"
	IssuingCardSpendingControlsSpendingLimitCategoryTelecommunicationServices                                  IssuingCardSpendingControlsSpendingLimitCategory = "telecommunication_services"
	IssuingCardSpendingControlsSpendingLimitCategoryTelegraphServices                                          IssuingCardSpendingControlsSpendingLimitCategory = "telegraph_services"
	IssuingCardSpendingControlsSpendingLimitCategoryTentAndAwningShops                                         IssuingCardSpendingControlsSpendingLimitCategory = "tent_and_awning_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryTestingLaboratories                                        IssuingCardSpendingControlsSpendingLimitCategory = "testing_laboratories"
	IssuingCardSpendingControlsSpendingLimitCategoryTheatricalTicketAgencies                                   IssuingCardSpendingControlsSpendingLimitCategory = "theatrical_ticket_agencies"
	IssuingCardSpendingControlsSpendingLimitCategoryTimeshares                                                 IssuingCardSpendingControlsSpendingLimitCategory = "timeshares"
	IssuingCardSpendingControlsSpendingLimitCategoryTireRetreadingAndRepair                                    IssuingCardSpendingControlsSpendingLimitCategory = "tire_retreading_and_repair"
	IssuingCardSpendingControlsSpendingLimitCategoryTollsBridgeFees                                            IssuingCardSpendingControlsSpendingLimitCategory = "tolls_bridge_fees"
	IssuingCardSpendingControlsSpendingLimitCategoryTouristAttractionsAndExhibits                              IssuingCardSpendingControlsSpendingLimitCategory = "tourist_attractions_and_exhibits"
	IssuingCardSpendingControlsSpendingLimitCategoryTowingServices                                             IssuingCardSpendingControlsSpendingLimitCategory = "towing_services"
	IssuingCardSpendingControlsSpendingLimitCategoryTrailerParksCampgrounds                                    IssuingCardSpendingControlsSpendingLimitCategory = "trailer_parks_campgrounds"
	IssuingCardSpendingControlsSpendingLimitCategoryTransportationServices                                     IssuingCardSpendingControlsSpendingLimitCategory = "transportation_services"
	IssuingCardSpendingControlsSpendingLimitCategoryTravelAgenciesTourOperators                                IssuingCardSpendingControlsSpendingLimitCategory = "travel_agencies_tour_operators"
	IssuingCardSpendingControlsSpendingLimitCategoryTruckStopIteration                                         IssuingCardSpendingControlsSpendingLimitCategory = "truck_stop_iteration"
	IssuingCardSpendingControlsSpendingLimitCategoryTruckUtilityTrailerRentals                                 IssuingCardSpendingControlsSpendingLimitCategory = "truck_utility_trailer_rentals"
	IssuingCardSpendingControlsSpendingLimitCategoryTypesettingPlateMakingAndRelatedServices                   IssuingCardSpendingControlsSpendingLimitCategory = "typesetting_plate_making_and_related_services"
	IssuingCardSpendingControlsSpendingLimitCategoryTypewriterStores                                           IssuingCardSpendingControlsSpendingLimitCategory = "typewriter_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryUSFederalGovernmentAgenciesOrDepartments                   IssuingCardSpendingControlsSpendingLimitCategory = "u_s_federal_government_agencies_or_departments"
	IssuingCardSpendingControlsSpendingLimitCategoryUniformsCommercialClothing                                 IssuingCardSpendingControlsSpendingLimitCategory = "uniforms_commercial_clothing"
	IssuingCardSpendingControlsSpendingLimitCategoryUsedMerchandiseAndSecondhandStores                         IssuingCardSpendingControlsSpendingLimitCategory = "used_merchandise_and_secondhand_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryUtilities                                                  IssuingCardSpendingControlsSpendingLimitCategory = "utilities"
	IssuingCardSpendingControlsSpendingLimitCategoryVarietyStores                                              IssuingCardSpendingControlsSpendingLimitCategory = "variety_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryVeterinaryServices                                         IssuingCardSpendingControlsSpendingLimitCategory = "veterinary_services"
	IssuingCardSpendingControlsSpendingLimitCategoryVideoAmusementGameSupplies                                 IssuingCardSpendingControlsSpendingLimitCategory = "video_amusement_game_supplies"
	IssuingCardSpendingControlsSpendingLimitCategoryVideoGameArcades                                           IssuingCardSpendingControlsSpendingLimitCategory = "video_game_arcades"
	IssuingCardSpendingControlsSpendingLimitCategoryVideoTapeRentalStores                                      IssuingCardSpendingControlsSpendingLimitCategory = "video_tape_rental_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryVocationalTradeSchools                                     IssuingCardSpendingControlsSpendingLimitCategory = "vocational_trade_schools"
	IssuingCardSpendingControlsSpendingLimitCategoryWatchJewelryRepair                                         IssuingCardSpendingControlsSpendingLimitCategory = "watch_jewelry_repair"
	IssuingCardSpendingControlsSpendingLimitCategoryWeldingRepair                                              IssuingCardSpendingControlsSpendingLimitCategory = "welding_repair"
	IssuingCardSpendingControlsSpendingLimitCategoryWholesaleClubs                                             IssuingCardSpendingControlsSpendingLimitCategory = "wholesale_clubs"
	IssuingCardSpendingControlsSpendingLimitCategoryWigAndToupeeStores                                         IssuingCardSpendingControlsSpendingLimitCategory = "wig_and_toupee_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryWiresMoneyOrders                                           IssuingCardSpendingControlsSpendingLimitCategory = "wires_money_orders"
	IssuingCardSpendingControlsSpendingLimitCategoryWomensAccessoryAndSpecialtyShops                           IssuingCardSpendingControlsSpendingLimitCategory = "womens_accessory_and_specialty_shops"
	IssuingCardSpendingControlsSpendingLimitCategoryWomensReadyToWearStores                                    IssuingCardSpendingControlsSpendingLimitCategory = "womens_ready_to_wear_stores"
	IssuingCardSpendingControlsSpendingLimitCategoryWreckingAndSalvageYards                                    IssuingCardSpendingControlsSpendingLimitCategory = "wrecking_and_salvage_yards"
)

// IssuingCardSpendingControlsSpendingLimitInterval is the list of possible values for the interval
// for a spending limit on an issuing card.
type IssuingCardSpendingControlsSpendingLimitInterval string

// List of values that IssuingCardShippingStatus can take.
const (
	IssuingCardSpendingControlsSpendingLimitIntervalAllTime          IssuingCardSpendingControlsSpendingLimitInterval = "all_time"
	IssuingCardSpendingControlsSpendingLimitIntervalDaily            IssuingCardSpendingControlsSpendingLimitInterval = "daily"
	IssuingCardSpendingControlsSpendingLimitIntervalMonthly          IssuingCardSpendingControlsSpendingLimitInterval = "monthly"
	IssuingCardSpendingControlsSpendingLimitIntervalPerAuthorization IssuingCardSpendingControlsSpendingLimitInterval = "per_authorization"
	IssuingCardSpendingControlsSpendingLimitIntervalWeekly           IssuingCardSpendingControlsSpendingLimitInterval = "weekly"
	IssuingCardSpendingControlsSpendingLimitIntervalYearly           IssuingCardSpendingControlsSpendingLimitInterval = "yearly"
)

// IssuingCardStatus is the list of possible values for status on an issuing card.
type IssuingCardStatus string

// List of values that IssuingCardStatus can take.
const (
	IssuingCardStatusActive   IssuingCardStatus = "active"
	IssuingCardStatusCanceled IssuingCardStatus = "canceled"
	IssuingCardStatusInactive IssuingCardStatus = "inactive"
)

// IssuingCardType is the type of an issuing card.
type IssuingCardType string

// List of values that IssuingCardType can take.
const (
	IssuingCardTypePhysical IssuingCardType = "physical"
	IssuingCardTypeVirtual  IssuingCardType = "virtual"
)

// IssuingCardShippingParams is the set of parameters that can be used for the shipping parameter.
type IssuingCardShippingParams struct {
	Address *AddressParams `form:"address"`
	Name    *string        `form:"name"`
	Service *string        `form:"service"`
	Type    *string        `form:"type"`
}

// IssuingCardSpendingControlsSpendingLimitParams is the set of parameters that can be used to
// represent a given spending limit for an issuing card.
type IssuingCardSpendingControlsSpendingLimitParams struct {
	Amount     *int64    `form:"amount"`
	Categories []*string `form:"categories"`
	Interval   *string   `form:"interval"`
}

// IssuingCardSpendingControlsParams is the set of parameters that can be used to configure
// the spending controls for an issuing card
type IssuingCardSpendingControlsParams struct {
	AllowedCategories []*string                                         `form:"allowed_categories"`
	BlockedCategories []*string                                         `form:"blocked_categories"`
	SpendingLimits    []*IssuingCardSpendingControlsSpendingLimitParams `form:"spending_limits"`
}

// IssuingCardParams is the set of parameters that can be used when creating or updating an issuing card.
type IssuingCardParams struct {
	Params             `form:"*"`
	CancellationReason *string                            `form:"cancellation_reason"`
	Cardholder         *string                            `form:"cardholder"`
	Currency           *string                            `form:"currency"`
	ReplacementFor     *string                            `form:"replacement_for"`
	ReplacementReason  *string                            `form:"replacement_reason"`
	Shipping           *IssuingCardShippingParams         `form:"shipping"`
	SpendingControls   *IssuingCardSpendingControlsParams `form:"spending_controls"`
	Status             *string                            `form:"status"`
	Type               *string                            `form:"type"`
}

// IssuingCardListParams is the set of parameters that can be used when listing issuing cards.
type IssuingCardListParams struct {
	ListParams   `form:"*"`
	Cardholder   *string           `form:"cardholder"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	ExpMonth     *int64            `form:"exp_month"`
	ExpYear      *int64            `form:"exp_year"`
	Last4        *string           `form:"last4"`
	Status       *string           `form:"status"`
	Type         *string           `form:"type"`
}

// IssuingCardShipping is the resource representing shipping on an issuing card.
type IssuingCardShipping struct {
	Address        *Address                   `json:"address"`
	Carrier        IssuingCardShippingCarrier `json:"carrier"`
	Eta            int64                      `json:"eta"`
	Name           string                     `json:"name"`
	Service        IssuingCardShippingService `json:"service"`
	Status         IssuingCardShippingStatus  `json:"status"`
	TrackingNumber string                     `json:"tracking_number"`
	TrackingURL    string                     `json:"tracking_url"`
	Type           IssuingCardShippingType    `json:"type"`
}

// IssuingCardSpendingControlsSpendingLimit is the resource representing a spending limit
// for an issuing card.
type IssuingCardSpendingControlsSpendingLimit struct {
	Amount     int64                                              `json:"amount"`
	Categories []IssuingCardSpendingControlsSpendingLimitCategory `json:"categories"`
	Interval   IssuingCardSpendingControlsSpendingLimitInterval   `json:"interval"`
}

// IssuingCardSpendingControls is the resource representing spending controls
// for an issuing card.
type IssuingCardSpendingControls struct {
	AllowedCategories      []IssuingCardSpendingControlsAllowedCategory `json:"allowed_categories"`
	BlockedCategories      []IssuingCardSpendingControlsBlockedCategory `json:"blocked_categories"`
	SpendingLimits         []*IssuingCardSpendingControlsSpendingLimit  `json:"spending_limits"`
	SpendingLimitsCurrency Currency                                     `json:"spending_limits_currency"`
}

// IssuingCard is the resource representing a Stripe issuing card.
type IssuingCard struct {
	APIResource
	Brand              string                        `json:"brand"`
	CancellationReason IssuingCardCancellationReason `json:"cancellation_reason"`
	Cardholder         *IssuingCardholder            `json:"cardholder"`
	Created            int64                         `json:"created"`
	Currency           Currency                      `json:"currency"`
	CVC                string                        `json:"cvc"`
	ExpMonth           int64                         `json:"exp_month"`
	ExpYear            int64                         `json:"exp_year"`
	ID                 string                        `json:"id"`
	Last4              string                        `json:"last4"`
	Livemode           bool                          `json:"livemode"`
	Metadata           map[string]string             `json:"metadata"`
	Number             string                        `json:"number"`
	Object             string                        `json:"object"`
	ReplacedBy         *IssuingCard                  `json:"replaced_by"`
	ReplacementFor     *IssuingCard                  `json:"replacement_for"`
	ReplacementReason  IssuingCardReplacementReason  `json:"replacement_reason"`
	Shipping           *IssuingCardShipping          `json:"shipping"`
	SpendingControls   *IssuingCardSpendingControls  `json:"spending_controls"`
	Status             IssuingCardStatus             `json:"status"`
	Type               IssuingCardType               `json:"type"`
}

// IssuingCardList is a list of issuing cards as retrieved from a list endpoint.
type IssuingCardList struct {
	APIResource
	ListMeta
	Data []*IssuingCard `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingCard.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingCard) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingCard IssuingCard
	var v issuingCard
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingCard(v)
	return nil
}
