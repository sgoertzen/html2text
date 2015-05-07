package html2text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextify(t *testing.T) {
	expected := "body\nbody2"
	r := Textify("<html><body><b>body</b><br/>body2</body></html>")
	assert.Equal(t, expected, r)
}

func TestBuildDescriptionDiv(t *testing.T) {
	expected := "first\nsecond"
	r := Textify("<div>first</div>second")
	assert.Equal(t, expected, r)
}

func TestBuildDescriptionLink(t *testing.T) {
	expected := "somelink (link: someurl)"
	r := Textify("<a href=\"someurl\">somelink</a>")
	assert.Equal(t, expected, r)
}

func TestBuildDescriptionSpaces(t *testing.T) {
	expected := "hello"
	r := Textify("<div> hello  </div>")
	assert.Equal(t, expected, r)
}

func TestBuildDescriptionMisc(t *testing.T) {
	expected := "a   a"
	r := Textify("a &nbsp; a")
	assert.Equal(t, expected, r)
}

func TestBuildDescriptionTable(t *testing.T) {
	expected := `Join by phone 
  1-877-668-4490 Call-in toll-free number (US/Canada) 
  1-408-792-6300 Call-in toll number (US/Canada) 
  Access code: 578 275 982 
  Global call-in numbers (link: https://akqa.webex.com/akqa/globalcallin.php?serviceType=MC&ED=299778282&tollFree=1)  |  Toll-free calling restrictions (link: http://www.webex.com/pdf/tollfree_restrictions.pdf)`

   test := `<table width="747" style="width:448.2pt;"> <col width="747" style="width:448.2pt;"> <tbody> <tr> <td><font face="Arial" color="#666666"><b>Join by phone</b></font></td> </tr> <tr> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;"><b>1-877-668-4490</b> Call-in toll-free number (US/Canada)</span></font></td> </tr> <tr> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;"><b>1-408-792-6300</b> Call-in toll number (US/Canada)</span></font></td> </tr> <tr> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;">Access code: 578 275 982</span></font></td> </tr> <tr> <td><a href="https://akqa.webex.com/akqa/globalcallin.php?serviceType=MC&amp;ED=299778282&amp;tollFree=1"><font face="Arial" size="2" color="#00AFF9"><span style="font-size:10pt;"><u>Global call-in numbers</u></span></font></a><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;">&nbsp; |&nbsp; </span></font><a href="http://www.webex.com/pdf/tollfree_restrictions.pdf"><font face="Arial" size="2" color="#00AFF9"><span style="font-size:10pt;"><u>Toll-free calling restrictions</u></span></font></a></td> </tr> </tbody> </table>`

	r := Textify(test)
	assert.Equal(t, expected, r)
}

func TestBuildDescriptionComment(t *testing.T) {
	expected := "this should appear"
	r := Textify("<!-- this should not appear -->this should appear")
	assert.Equal(t, expected, r)
}

func TestBuildDescriptionCommentInHead(t *testing.T) {
	expected := "qwerty"
	
   body := `<html> <head> <meta http-equiv="Content-Type" content="text/html; charset=utf-8"> <meta name="Generator" content="Microsoft Exchange Server"> <!-- converted from rtf --><style><!-- .EmailQuote { margin-left: 1pt; padding-left: 4pt; border-left: #800000 2px solid; } --></style> </head> <body>qwerty</body> </html>`

	r := Textify(body)
	assert.Equal(t, expected, r)
}


func TestBuildDescriptionComplex(t *testing.T) {
	expected := `Meeting Agenda:
  
   Quick program/project statuses by each AKQA product manager
Sprint development prioritization (if needed)
CODE/AKQA resource alignment
Content and Creative team prioritization 
  
     Join WebEx meeting (link: https://akqa.webex.com/akqa/j.php?MTID=m54a67710a112cf53d262d8b5f7acdf66) 
    
     Meeting number: 578 275 982 
  Meeting password: audiusa2015 
    
      
    
     Join by phone 
  1-877-668-4490 Call-in toll-free number (US/Canada) 
  1-408-792-6300 Call-in toll number (US/Canada) 
  Access code: 578 275 982 
  Global call-in numbers (link: https://akqa.webex.com/akqa/globalcallin.php?serviceType=MC&ED=299778282&tollFree=1)  |  Toll-free calling restrictions (link: http://www.webex.com/pdf/tollfree_restrictions.pdf)`
	
   body := `<html> <head> <meta http-equiv="Content-Type" content="text/html; charset=utf-8"> <meta name="Generator" content="Microsoft Exchange Server"> <!-- converted from rtf --><style><!-- .EmailQuote { margin-left: 1pt; padding-left: 4pt; border-left: #800000 2px solid; } --></style> </head> <body> <font face="Times New Roman" size="3"><span style="font-size:12pt;"><a name="BM_BEGIN"></a> <div><font face="Calibri" size="2"><span style="font-size:11pt;">Meeting Agenda:</span></font></div> <div><font face="Calibri" size="2"><span style="font-size:11pt;">&nbsp;</span></font></div> <ol style="margin:0;padding-left:30pt;"> <font face="Calibri" size="2"><span style="font-size:11pt;"> <li>Quick program/project statuses by each AKQA product manager</li><li>Sprint development prioritization (if needed)</li><li>CODE/AKQA resource alignment</li><li>Content and Creative team prioritization<a name="BM__InsertRtfSavedPosition"></a></span></font> </li></ol> <div><font face="Calibri" size="2"><span style="font-size:11pt;">&nbsp;</span></font></div> <table width="747" style="width:448.2pt;"> <col width="747" style="width:448.2pt;"> <tbody> <tr> <td><a href="https://akqa.webex.com/akqa/j.php?MTID=m54a67710a112cf53d262d8b5f7acdf66"><font face="Arial" color="#00AFF9"><b><u>Join WebEx meeting</u></b></font></a></td> </tr> </tbody> </table> <div><font face="Calibri" size="2"><span style="font-size:11pt;">&nbsp;</span></font></div> <table width="747" style="width:448.2pt;"> <col width="378" style="width:226.8pt;"><col width="369" style="width:221.4pt;"> <tbody> <tr> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;">Meeting number:</span></font></td> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;">578 275 982</span></font></td> </tr> <tr> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;">Meeting password:</span></font></td> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;">audiusa2015</span></font></td> </tr> </tbody> </table> <div><font face="Calibri" size="2"><span style="font-size:11pt;">&nbsp;</span></font></div> <table width="747" style="width:448.2pt;"> <col width="747" style="width:448.2pt;"> <tbody> <tr height="25" style="height:15pt;"> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;"></span></font></td> </tr> </tbody> </table> <div><font face="Calibri" size="2"><span style="font-size:11pt;">&nbsp;</span></font></div> <table width="747" style="width:448.2pt;"> <col width="747" style="width:448.2pt;"> <tbody> <tr> <td><font face="Arial" color="#666666"><b>Join by phone</b></font></td> </tr> <tr> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;"><b>1-877-668-4490</b> Call-in toll-free number (US/Canada)</span></font></td> </tr> <tr> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;"><b>1-408-792-6300</b> Call-in toll number (US/Canada)</span></font></td> </tr> <tr> <td><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;">Access code: 578 275 982</span></font></td> </tr> <tr> <td><a href="https://akqa.webex.com/akqa/globalcallin.php?serviceType=MC&amp;ED=299778282&amp;tollFree=1"><font face="Arial" size="2" color="#00AFF9"><span style="font-size:10pt;"><u>Global call-in numbers</u></span></font></a><font face="Arial" size="3" color="#666666"><span style="font-size:11.5pt;">&nbsp; |&nbsp; </span></font><a href="http://www.webex.com/pdf/tollfree_restrictions.pdf"><font face="Arial" size="2" color="#00AFF9"><span style="font-size:10pt;"><u>Toll-free calling restrictions</u></span></font></a></td> </tr> </tbody> </table> <div><font face="Calibri" size="2"><span style="font-size:11pt;">&nbsp;</span></font></div> <div><font face="Calibri" size="2"><span style="font-size:11pt;">&nbsp;</span></font></div> </span></font> </body> </html>`

	r := Textify(body)
	assert.Equal(t, expected, r)
}