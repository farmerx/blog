## 参考文档
http://officeopenxml.com/anatomyofOOXML-xlsx.php

## xlsx Package Structure

`SpreadsheetML`或`.xlsx`文件是包含许多“部件”（通常为UTF-8或UTF-16编码）或XML文件的zip文件（包）。该包还可能包含其他媒体文件，如图片。
该结构根据OOXML标准ECMA-376第2部分中概述的开放式包装惯例进行组织。


只需解压缩.xlsx文件，即可查看包含SpreadsheetML文件的文件结构和文件。

```
 - _rels  关系:每个包都包含一个关系部分，用于定义其他部分与包外部资源之间的关系。这将关系与内容分开，使得在不更改引用目标的源的情况下更改关系变得容易。
 - docProps
 - xl
 - [Content_Types].xml  内容类型: 每个包的根路径下都会有
```

部件的数量和类型将根据电子表格中的内容而有所不同，但始终会有[Content_Types] .xml，一个或多个关系部件，工作簿部件和至少一个工作表。电子表格的核心数据包含在工作表部分中，在xsxl内容概述中有更详细的讨论(http://officeopenxml.com/SScontentOverview.php)。

## [Content_Types].xml 

每个包都必须有一个[Content_Types] .xml，在包的根目录中找到。此文件包含包中所有部件的内容类型的列表。每个部分及其类型必须列在[Content_Types] .xml中。以下是主要内容部分的内容类型：
```
<Override PartName="/xl/workbook.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"/>
```
在将新零件添加到包装时，请务必牢记这一点。

## _rels

每个包都包含一个关系部分，用于定义其他部分与包外部资源之间的关系。这将关系与内容分开，使得在不更改引用目标的源的情况下更改关系变得容易。
```
- _rels
- - .rels
```
对于OOXML包，_rels文件夹中始终存在关系部分（.rels），用于标识包的起始部分或包关系。例如，以下定义内容的开始部分的标识：
```
<Relationship Id ="rId1" Type ="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target ="xl / workbook.xml"/>
```
对于app.xml和core.xml，.rels中通常也存在关系。

除了包的关系部分之外，作为一个或多个关系的源的每个部分将具有其自己的关系部分。每个这样的关系部分都在零件的_rels子文件夹中找到，并通过将“.rels”附加到零件的名称来命名。

通常，主要内容部分（workbook.xml）具有自己的关系部分（workbook.xml.rels）。它将包含与内容其他部分的关系，例如sheet1.xml，sharedStrings.xml，styles.xml，theme1，xml以及外部链接的URI。

```

```
关系可以是显式的，也可以是隐式的。对于显式关系，使用<Relationship>元素的Id属性引用资源 。即， 标识在源直接映射到标识的关系的项目，与显式引用到目标。

例如，工作表可能包含如下链接：
```
<w:hyperlink ref ="A11" r:id ="rId4">
```
在R:ID ="rId4"引用关系的一部分工作表（worksheet1.xml.rels）内的下列关系。
```
<Relationship Id="rId4" Type="http://../hyperlink" Target ="http://www.google.com/" TargetMode ="External"/>
```
对于隐式关系，没有对<Relationship> Id的直接引用。相反，参考是理解的。





