[![Go Report Card](https://goreportcard.com/badge/github.com/ip2whois/ip2who-go)](https://goreportcard.com/report/github.com/ip2whois/ip2whois-go)

IP2WHOIS Go Package
=======================
This Go package enables user to easily implement the checking of WHOIS information for a particular domain into their solution using the API from https://www.ip2whois.com. It is a WHOIS lookup api that helps users to obtain domain information, WHOIS record, by using a domain name. The WHOIS API returns a comprehensive WHOIS data such as creation date, updated date, expiration date, domain age, the contact information of the registrant, mailing address, phone number, email address, nameservers the domain is using and much more. IP2WHOIS supports the query for [1113 TLDs and 634 ccTLDs](https://www.ip2whois.com/tld-cctld-supported).

This package requires API key to function. You may sign up for a free API key at https://www.ip2whois.com/register.


Installation
============
To install this package type the following:

```
go get github.com/ip2whois/ip2whois-go
```


Usage Example
=============
### Lookup Domain Information
```go
package main

import (
	"github.com/ip2whois/ip2whois-go"
	"fmt"
)

func main() {
	apikey := "YOUR_API_KEY"
	
	api := ip2whois.OpenApi(apikey)
	
	domain := "example.com"
	res, err := api.LookUp(domain)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	fmt.Printf("Domain => %+v\n", res.Domain)
	fmt.Printf("DomainID => %+v\n", res.DomainID)
	fmt.Printf("Status => %+v\n", res.Status)
	fmt.Printf("CreateDate => %+v\n", res.CreateDate)
	fmt.Printf("ExpireDate => %+v\n", res.ExpireDate)
	fmt.Printf("DomainAge => %+v\n", res.DomainAge)
	fmt.Printf("WhoisServer => %+v\n", res.WhoisServer)
	
	fmt.Printf("Registrar.IanaID => %+v\n", res.Registrar.IanaID)
	fmt.Printf("Registrar.Name => %+v\n", res.Registrar.Name)
	fmt.Printf("Registrar.URL => %+v\n", res.Registrar.URL)
	
	fmt.Printf("Registrant.Name => %+v\n", res.Registrant.Name)
	fmt.Printf("Registrant.Organization => %+v\n", res.Registrant.Organization)
	fmt.Printf("Registrant.StreetAddress => %+v\n", res.Registrant.StreetAddress)
	fmt.Printf("Registrant.City => %+v\n", res.Registrant.City)
	fmt.Printf("Registrant.Region => %+v\n", res.Registrant.Region)
	fmt.Printf("Registrant.ZipCode => %+v\n", res.Registrant.ZipCode)
	fmt.Printf("Registrant.Country => %+v\n", res.Registrant.Country)
	fmt.Printf("Registrant.Phone => %+v\n", res.Registrant.Phone)
	fmt.Printf("Registrant.Fax => %+v\n", res.Registrant.Fax)
	fmt.Printf("Registrant.Email => %+v\n", res.Registrant.Email)
	
	fmt.Printf("Admin.Name => %+v\n", res.Admin.Name)
	fmt.Printf("Admin.Organization => %+v\n", res.Admin.Organization)
	fmt.Printf("Admin.StreetAddress => %+v\n", res.Admin.StreetAddress)
	fmt.Printf("Admin.City => %+v\n", res.Admin.City)
	fmt.Printf("Admin.Region => %+v\n", res.Admin.Region)
	fmt.Printf("Admin.ZipCode => %+v\n", res.Admin.ZipCode)
	fmt.Printf("Admin.Country => %+v\n", res.Admin.Country)
	fmt.Printf("Admin.Phone => %+v\n", res.Admin.Phone)
	fmt.Printf("Admin.Fax => %+v\n", res.Admin.Fax)
	fmt.Printf("Admin.Email => %+v\n", res.Admin.Email)
	
	fmt.Printf("Tech.Name => %+v\n", res.Tech.Name)
	fmt.Printf("Tech.Organization => %+v\n", res.Tech.Organization)
	fmt.Printf("Tech.StreetAddress => %+v\n", res.Tech.StreetAddress)
	fmt.Printf("Tech.City => %+v\n", res.Tech.City)
	fmt.Printf("Tech.Region => %+v\n", res.Tech.Region)
	fmt.Printf("Tech.ZipCode => %+v\n", res.Tech.ZipCode)
	fmt.Printf("Tech.Country => %+v\n", res.Tech.Country)
	fmt.Printf("Tech.Phone => %+v\n", res.Tech.Phone)
	fmt.Printf("Tech.Fax => %+v\n", res.Tech.Fax)
	fmt.Printf("Tech.Email => %+v\n", res.Tech.Email)
	
	fmt.Printf("Billing.Name => %+v\n", res.Billing.Name)
	fmt.Printf("Billing.Organization => %+v\n", res.Billing.Organization)
	fmt.Printf("Billing.StreetAddress => %+v\n", res.Billing.StreetAddress)
	fmt.Printf("Billing.City => %+v\n", res.Billing.City)
	fmt.Printf("Billing.Region => %+v\n", res.Billing.Region)
	fmt.Printf("Billing.ZipCode => %+v\n", res.Billing.ZipCode)
	fmt.Printf("Billing.Country => %+v\n", res.Billing.Country)
	fmt.Printf("Billing.Phone => %+v\n", res.Billing.Phone)
	fmt.Printf("Billing.Fax => %+v\n", res.Billing.Fax)
	fmt.Printf("Billing.Email => %+v\n", res.Billing.Email)
	
	fmt.Printf("Nameservers => %+v\n", res.Nameservers)
}
```

### Convert Normal Text to Punycode
```go
package main

import (
	"github.com/ip2whois/ip2whois-go"
	"fmt"
)

func main() {
	domain = "täst.de"
	res, err := api.GetPunycode(domain)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v => %+v\n", domain, res);
}
```

### Convert Punycode to Normal Text
```go
package main

import (
	"github.com/ip2whois/ip2whois-go"
	"fmt"
)

func main() {
	domain = "xn--tst-qla.de"
	res, err := api.GetNormalText(domain)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v => %+v\n", domain, res);
}
```


Response Parameter
==================
### Lookup function
| Parameter | Type | Description |
|---|---|---|
|domain|string|Domain name.|
|domain_id|string|Domain name ID.|
|status|string|Domain name status.|
|create_date|string|Domain name creation date.|
|update_date|string|Domain name updated date.|
|expire_date|string|Domain name expiration date.|
|domain_age|integer|Domain name age in day(s).|
|whois_server|string|WHOIS server name.|
|registrar.iana_id|string|Registrar IANA ID.|
|registrar.name|string|Registrar name.|
|registrar.url|string|Registrar URL.|
|registrant.name|string|Registrant name.|
|registrant.organization|string|Registrant organization.|
|registrant.street_address|string|Registrant street address.|
|registrant.city|string|Registrant city.|
|registrant.region|string|Registrant region.|
|registrant.zip_code|string|Registrant ZIP Code.|
|registrant.country|string|Registrant country.|
|registrant.phone|string|Registrant phone number.|
|registrant.fax|string|Registrant fax number.|
|registrant.email|string|Registrant email address.|
|admin.name|string|Admin name.|
|admin.organization|string|Admin organization.|
|admin.street_address|string|Admin street address.|
|admin.city|string|Admin city.|
|admin.region|string|Admin region.|
|admin.zip_code|string|Admin ZIP Code.|
|admin.country|string|Admin country.|
|admin.phone|string|Admin phone number.|
|admin.fax|string|Admin fax number.|
|admin.email|string|Admin email address.|
|tech.name|string|Tech name.|
|tech.organization|string|Tech organization.|
|tech.street_address|string|Tech street address.|
|tech.city|string|Tech city.|
|tech.region|string|Tech region.|
|tech.zip_code|string|Tech ZIP Code.|
|tech.country|string|Tech country.|
|tech.phone|string|Tech phone number.|
|tech.fax|string|Tech fax number.|
|tech.email|string|Tech email address.|
|billing.name|string|Billing name.|
|billing.organization|string|Billing organization.|
|billing.street_address|string|Billing street address.|
|billing.city|string|Billing city.|
|billing.region|string|Billing region.|
|billing.zip_code|string|Billing ZIP Code.|
|billing.country|string|Billing country.|
|billing.phone|string|Billing phone number.|
|billing.fax|string|Billing fax number.|
|billing.email|string|Billing email address.|
|name_servers|array|Name servers|

```json
{
    "domain": "greendot.com",
    "domain_id": "600750_DOMAIN_COM-VRSN",
    "status": "registered",
    "create_date": "1997-11-03T00:00:00Z",
    "update_date": "2019-10-29T01:25:57Z",
    "expire_date": "2021-11-02T05:00:00Z",
    "domain_age": 9027,
    "whois_server": "whois.corporatedomains.com",
    "registrar": {
        "iana_id": "299",
        "name": "CSC CORPORATE DOMAINS, INC.",
        "url": "www.cscprotectsbrands.com"
    },
    "registrant": {
        "name": "Admin Role",
        "organization": "Green Dot Corporation",
        "street_address": "",
        "city": "Pasadena",
        "region": "CA",
        "zip_code": "91107",
        "country": "US",
        "phone": "+1.8664120548",
        "fax": "+1.8664120548",
        "email": "adminrole@greendotcorp.com"
    },
    "admin": {
        "name": "Admin Role",
        "organization": "Green Dot Corporation",
        "street_address": "",
        "city": "Pasadena",
        "region": "CA",
        "zip_code": "91107",
        "country": "US",
        "phone": "+1.8664120548",
        "fax": "+1.8664120548",
        "email": "adminrole@greendotcorp.com"
    },
    "tech": {
        "name": "Admin Role",
        "organization": "Green Dot Corporation",
        "street_address": "",
        "city": "Pasadena",
        "region": "CA",
        "zip_code": "91107",
        "country": "US",
        "phone": "+1.8664120548",
        "fax": "+1.8664120548",
        "email": "adminrole@greendotcorp.com"
    },
    "billing": {
        "name": "",
        "organization": "",
        "street_address": "",
        "city": "",
        "region": "",
        "zip_code": "",
        "country": "",
        "phone": "",
        "fax": "",
        "email": ""
    },
    "nameservers": "ns1.greendotdns.com, ns2.greendotdns.com"
}
```


LICENCE
=====================
See the LICENSE file.
