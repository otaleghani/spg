// /******************************************************************
// 
// Author      Oliviero Taleghani
// Date        2024-05-14
// 
// Description
// Creator of emails.
// 
// Usage
// 
// Dependency
//   "github.com/otaleghani/spg/internal/parser"
//   "github.com/otaleghani/spg/internal/formatter"
// 
// Todo
// 
// Changelog
//   [0.0.1]   2024-05-13
//   Added     Initial release
// 
// ********************************************************************/

package spg
// 
// // How is an email created
// //
// // name, domain
// // 
// // Different algorithmis to create an email
// 
// 
// // Il generator dovrà contenere tipologia algo per nome, e tipologia
// // algo per domain.
// 
// type EmailGenerator struct {
//   Options GeneratorMetadata
//   UsernameAlgo string
//   DomainAlgo string
//   Data_FirstNames []string
//   Data_LastNames []string
//   Data_Words []string
// }
// 
// func Email(
//   lang string,
//   userAlgo string,
//   domainAlgo string,
// ) (EmailGenerator, error) {
//   words, err := Words(lang, "lower")
//   if err != nil {
//     return EmailGenerator{}, err
//   }
//   firstNames, err := FirstNames(lang, format)
//   if err != nil {
//     return EmailGenerator{}, err
//   }
//   lastNames, err := LastNames(lang, format)
//   if err != nil {
//     return EmailGenerator{}, err
//   }
//   err := verifyUsernameAlgo()
//   if err != nil {
//     return EmailGenerator{}, err
//   }
//   err := verifyDomainAlgo()
//   if err != nil {
//     return EmailGenerator{}, err
//   }
//   result := EmailGenerator{
//     Options: GeneratorMetadata{
//       Lang: lang,
//       Format: lowerFormat,
//     },
//     UsernameAlgo: userAlgo,
//     DomainAlgo: domainAlgo,
//     Data_FirstNames: firstNames.Data,
//     Data_LastNames: lastNames.Data,
//     Data_Words: words.Data,
//   }
//   return result, nil
// }
// 
// func (gen EmailGenerator) Get() string {
//   username, domain := "", ""
//   switch gen.UsernameAlgo {
//   case "f.lastname":
//     Generator{Data:gen.Data_FirstNames}.Get()
//     // first char of first name
//     // first char of last name
//     //
//   }
// 
//   return username + "@" + domain, nil 
// }
