/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"time"
)

// checks if the SubscriptionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDTO{}

// SubscriptionDTO Represents the Subscription data exposed to the API
type SubscriptionDTO struct {
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Owner *NestedOwnerDTO `json:"owner,omitempty"`
	Product *ProductDTO `json:"product,omitempty"`
	DerivedProduct *ProductDTO `json:"derivedProduct,omitempty"`
	ProvidedProducts []ProductDTO `json:"providedProducts,omitempty"`
	DerivedProvidedProducts []ProductDTO `json:"derivedProvidedProducts,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	ContractNumber *string `json:"contractNumber,omitempty"`
	AccountNumber *string `json:"accountNumber,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
	LastModified *time.Time `json:"lastModified,omitempty"`
	OrderNumber *string `json:"orderNumber,omitempty"`
	UpstreamPoolId *string `json:"upstreamPoolId,omitempty"`
	UpstreamEntitlementId *string `json:"upstreamEntitlementId,omitempty"`
	UpstreamConsumerId *string `json:"upstreamConsumerId,omitempty"`
	Certificate *CertificateDTO `json:"certificate,omitempty"`
	Cdn *CdnDTO `json:"cdn,omitempty"`
	Stacked *bool `json:"stacked,omitempty"`
	StackId *string `json:"stackId,omitempty"`
}

// NewSubscriptionDTO instantiates a new SubscriptionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDTO() *SubscriptionDTO {
	this := SubscriptionDTO{}
	return &this
}

// NewSubscriptionDTOWithDefaults instantiates a new SubscriptionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDTOWithDefaults() *SubscriptionDTO {
	this := SubscriptionDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SubscriptionDTO) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *SubscriptionDTO) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionDTO) SetId(v string) {
	o.Id = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetOwner() NestedOwnerDTO {
	if o == nil || IsNil(o.Owner) {
		var ret NestedOwnerDTO
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetOwnerOk() (*NestedOwnerDTO, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NestedOwnerDTO and assigns it to the Owner field.
func (o *SubscriptionDTO) SetOwner(v NestedOwnerDTO) {
	o.Owner = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetProduct() ProductDTO {
	if o == nil || IsNil(o.Product) {
		var ret ProductDTO
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetProductOk() (*ProductDTO, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ProductDTO and assigns it to the Product field.
func (o *SubscriptionDTO) SetProduct(v ProductDTO) {
	o.Product = &v
}

// GetDerivedProduct returns the DerivedProduct field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetDerivedProduct() ProductDTO {
	if o == nil || IsNil(o.DerivedProduct) {
		var ret ProductDTO
		return ret
	}
	return *o.DerivedProduct
}

// GetDerivedProductOk returns a tuple with the DerivedProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetDerivedProductOk() (*ProductDTO, bool) {
	if o == nil || IsNil(o.DerivedProduct) {
		return nil, false
	}
	return o.DerivedProduct, true
}

// HasDerivedProduct returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasDerivedProduct() bool {
	if o != nil && !IsNil(o.DerivedProduct) {
		return true
	}

	return false
}

// SetDerivedProduct gets a reference to the given ProductDTO and assigns it to the DerivedProduct field.
func (o *SubscriptionDTO) SetDerivedProduct(v ProductDTO) {
	o.DerivedProduct = &v
}

// GetProvidedProducts returns the ProvidedProducts field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetProvidedProducts() []ProductDTO {
	if o == nil || IsNil(o.ProvidedProducts) {
		var ret []ProductDTO
		return ret
	}
	return o.ProvidedProducts
}

// GetProvidedProductsOk returns a tuple with the ProvidedProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetProvidedProductsOk() ([]ProductDTO, bool) {
	if o == nil || IsNil(o.ProvidedProducts) {
		return nil, false
	}
	return o.ProvidedProducts, true
}

// HasProvidedProducts returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasProvidedProducts() bool {
	if o != nil && !IsNil(o.ProvidedProducts) {
		return true
	}

	return false
}

// SetProvidedProducts gets a reference to the given []ProductDTO and assigns it to the ProvidedProducts field.
func (o *SubscriptionDTO) SetProvidedProducts(v []ProductDTO) {
	o.ProvidedProducts = v
}

// GetDerivedProvidedProducts returns the DerivedProvidedProducts field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetDerivedProvidedProducts() []ProductDTO {
	if o == nil || IsNil(o.DerivedProvidedProducts) {
		var ret []ProductDTO
		return ret
	}
	return o.DerivedProvidedProducts
}

// GetDerivedProvidedProductsOk returns a tuple with the DerivedProvidedProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetDerivedProvidedProductsOk() ([]ProductDTO, bool) {
	if o == nil || IsNil(o.DerivedProvidedProducts) {
		return nil, false
	}
	return o.DerivedProvidedProducts, true
}

// HasDerivedProvidedProducts returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasDerivedProvidedProducts() bool {
	if o != nil && !IsNil(o.DerivedProvidedProducts) {
		return true
	}

	return false
}

// SetDerivedProvidedProducts gets a reference to the given []ProductDTO and assigns it to the DerivedProvidedProducts field.
func (o *SubscriptionDTO) SetDerivedProvidedProducts(v []ProductDTO) {
	o.DerivedProvidedProducts = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *SubscriptionDTO) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *SubscriptionDTO) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *SubscriptionDTO) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetContractNumber returns the ContractNumber field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetContractNumber() string {
	if o == nil || IsNil(o.ContractNumber) {
		var ret string
		return ret
	}
	return *o.ContractNumber
}

// GetContractNumberOk returns a tuple with the ContractNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetContractNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ContractNumber) {
		return nil, false
	}
	return o.ContractNumber, true
}

// HasContractNumber returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasContractNumber() bool {
	if o != nil && !IsNil(o.ContractNumber) {
		return true
	}

	return false
}

// SetContractNumber gets a reference to the given string and assigns it to the ContractNumber field.
func (o *SubscriptionDTO) SetContractNumber(v string) {
	o.ContractNumber = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *SubscriptionDTO) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *SubscriptionDTO) SetModified(v time.Time) {
	o.Modified = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *SubscriptionDTO) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetOrderNumber() string {
	if o == nil || IsNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNumber) {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasOrderNumber() bool {
	if o != nil && !IsNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *SubscriptionDTO) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetUpstreamPoolId returns the UpstreamPoolId field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetUpstreamPoolId() string {
	if o == nil || IsNil(o.UpstreamPoolId) {
		var ret string
		return ret
	}
	return *o.UpstreamPoolId
}

// GetUpstreamPoolIdOk returns a tuple with the UpstreamPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetUpstreamPoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamPoolId) {
		return nil, false
	}
	return o.UpstreamPoolId, true
}

// HasUpstreamPoolId returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasUpstreamPoolId() bool {
	if o != nil && !IsNil(o.UpstreamPoolId) {
		return true
	}

	return false
}

// SetUpstreamPoolId gets a reference to the given string and assigns it to the UpstreamPoolId field.
func (o *SubscriptionDTO) SetUpstreamPoolId(v string) {
	o.UpstreamPoolId = &v
}

// GetUpstreamEntitlementId returns the UpstreamEntitlementId field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetUpstreamEntitlementId() string {
	if o == nil || IsNil(o.UpstreamEntitlementId) {
		var ret string
		return ret
	}
	return *o.UpstreamEntitlementId
}

// GetUpstreamEntitlementIdOk returns a tuple with the UpstreamEntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetUpstreamEntitlementIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamEntitlementId) {
		return nil, false
	}
	return o.UpstreamEntitlementId, true
}

// HasUpstreamEntitlementId returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasUpstreamEntitlementId() bool {
	if o != nil && !IsNil(o.UpstreamEntitlementId) {
		return true
	}

	return false
}

// SetUpstreamEntitlementId gets a reference to the given string and assigns it to the UpstreamEntitlementId field.
func (o *SubscriptionDTO) SetUpstreamEntitlementId(v string) {
	o.UpstreamEntitlementId = &v
}

// GetUpstreamConsumerId returns the UpstreamConsumerId field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetUpstreamConsumerId() string {
	if o == nil || IsNil(o.UpstreamConsumerId) {
		var ret string
		return ret
	}
	return *o.UpstreamConsumerId
}

// GetUpstreamConsumerIdOk returns a tuple with the UpstreamConsumerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetUpstreamConsumerIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamConsumerId) {
		return nil, false
	}
	return o.UpstreamConsumerId, true
}

// HasUpstreamConsumerId returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasUpstreamConsumerId() bool {
	if o != nil && !IsNil(o.UpstreamConsumerId) {
		return true
	}

	return false
}

// SetUpstreamConsumerId gets a reference to the given string and assigns it to the UpstreamConsumerId field.
func (o *SubscriptionDTO) SetUpstreamConsumerId(v string) {
	o.UpstreamConsumerId = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetCertificate() CertificateDTO {
	if o == nil || IsNil(o.Certificate) {
		var ret CertificateDTO
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetCertificateOk() (*CertificateDTO, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given CertificateDTO and assigns it to the Certificate field.
func (o *SubscriptionDTO) SetCertificate(v CertificateDTO) {
	o.Certificate = &v
}

// GetCdn returns the Cdn field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetCdn() CdnDTO {
	if o == nil || IsNil(o.Cdn) {
		var ret CdnDTO
		return ret
	}
	return *o.Cdn
}

// GetCdnOk returns a tuple with the Cdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetCdnOk() (*CdnDTO, bool) {
	if o == nil || IsNil(o.Cdn) {
		return nil, false
	}
	return o.Cdn, true
}

// HasCdn returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasCdn() bool {
	if o != nil && !IsNil(o.Cdn) {
		return true
	}

	return false
}

// SetCdn gets a reference to the given CdnDTO and assigns it to the Cdn field.
func (o *SubscriptionDTO) SetCdn(v CdnDTO) {
	o.Cdn = &v
}

// GetStacked returns the Stacked field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetStacked() bool {
	if o == nil || IsNil(o.Stacked) {
		var ret bool
		return ret
	}
	return *o.Stacked
}

// GetStackedOk returns a tuple with the Stacked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetStackedOk() (*bool, bool) {
	if o == nil || IsNil(o.Stacked) {
		return nil, false
	}
	return o.Stacked, true
}

// HasStacked returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasStacked() bool {
	if o != nil && !IsNil(o.Stacked) {
		return true
	}

	return false
}

// SetStacked gets a reference to the given bool and assigns it to the Stacked field.
func (o *SubscriptionDTO) SetStacked(v bool) {
	o.Stacked = &v
}

// GetStackId returns the StackId field value if set, zero value otherwise.
func (o *SubscriptionDTO) GetStackId() string {
	if o == nil || IsNil(o.StackId) {
		var ret string
		return ret
	}
	return *o.StackId
}

// GetStackIdOk returns a tuple with the StackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDTO) GetStackIdOk() (*string, bool) {
	if o == nil || IsNil(o.StackId) {
		return nil, false
	}
	return o.StackId, true
}

// HasStackId returns a boolean if a field has been set.
func (o *SubscriptionDTO) HasStackId() bool {
	if o != nil && !IsNil(o.StackId) {
		return true
	}

	return false
}

// SetStackId gets a reference to the given string and assigns it to the StackId field.
func (o *SubscriptionDTO) SetStackId(v string) {
	o.StackId = &v
}

func (o SubscriptionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.DerivedProduct) {
		toSerialize["derivedProduct"] = o.DerivedProduct
	}
	if !IsNil(o.ProvidedProducts) {
		toSerialize["providedProducts"] = o.ProvidedProducts
	}
	if !IsNil(o.DerivedProvidedProducts) {
		toSerialize["derivedProvidedProducts"] = o.DerivedProvidedProducts
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.ContractNumber) {
		toSerialize["contractNumber"] = o.ContractNumber
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if !IsNil(o.UpstreamPoolId) {
		toSerialize["upstreamPoolId"] = o.UpstreamPoolId
	}
	if !IsNil(o.UpstreamEntitlementId) {
		toSerialize["upstreamEntitlementId"] = o.UpstreamEntitlementId
	}
	if !IsNil(o.UpstreamConsumerId) {
		toSerialize["upstreamConsumerId"] = o.UpstreamConsumerId
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.Cdn) {
		toSerialize["cdn"] = o.Cdn
	}
	if !IsNil(o.Stacked) {
		toSerialize["stacked"] = o.Stacked
	}
	if !IsNil(o.StackId) {
		toSerialize["stackId"] = o.StackId
	}
	return toSerialize, nil
}

type NullableSubscriptionDTO struct {
	value *SubscriptionDTO
	isSet bool
}

func (v NullableSubscriptionDTO) Get() *SubscriptionDTO {
	return v.value
}

func (v *NullableSubscriptionDTO) Set(val *SubscriptionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDTO(val *SubscriptionDTO) *NullableSubscriptionDTO {
	return &NullableSubscriptionDTO{value: val, isSet: true}
}

func (v NullableSubscriptionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


