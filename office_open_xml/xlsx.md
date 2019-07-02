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

## SpreadsheetML文档特有的部分

下面列出了特定于SpreadsheetML电子表格的SpreadsheetML包的可能部分。请注意，电子表格可能只包含其中的一部分。例如，如果电子表格没有数据透视表，则数据包中不会包含数据透视表部件。

| 部分   |      描述   |
|----------|:-------------:|
|计算链(Calculation Chain) |当从公式计算单元格的值时，计算的顺序可能受计算值的顺序的影响。此部分包含指定顺序。包只能包含一个这样的部分。 |
|Chartsheet |包含存储在其自身工作表中的图表。包可以包含多个这样的部分，从工作簿部分引用。   | 
|评论(Comments)|包含给定工作表的注释。由于可能有多个工作表，因此可能会有多个注释部分。|
|连接(Connections)|电子表格可能与外部数据源有连接。此部分包含此类连接，解释了如何获取此类外部数据以及如何在工作簿中使用连接。可能只有一个这样的部分|
|自定义属性(Custom Property)|包含用户定义的数据。可能存在零个或多个这样的部分。|
|客户XML映射(Customer XML Mappings)|包含XML文件的架构，以及允许将架构映射到电子表格时要使用的行为的信息。可能只有一个这样的部分。|
|Dialogsheet|包含有关用户表单的旧客户对话框的信息。可能存在零个或多个这样的部分。|
|图纸(Drawings)|包含工作表中存在的一个或多个绘图元素的表示和布局信息。每个工作表都应该有绘图部分。|
|外部工作簿参考(External Workbook References)|包含有关其他电子表格包中引用的数据的信息。例如，电子表格可能有一个单元格，其值是根据另一个电子表格中的数据计算的。可能存在零个或多个这样的部分。|
|元数据(Metadata)|包含与通过在线分析处理（OLAP）技术与一个或多个其他单元格相关的单元格相关的信息。|
|数据透视表(Pivot Table)|包含数据透视表的定义。它描述了数据透视表布局的细节，指示了行轴，列轴和数据透视表值区域的字段。它表示数据透视表的格式。包中的每个数据透视表都有一个数据透视表部件。|
|数据透视表缓存定义(Pivot Table Cache Definition)|数据透视表缓存定义定义了数据透视缓存记录部分中的每个字段（即基础数据），包括字段名称和有关字段中包含的数据的信息。包中的每个数据透视表都有一个数据透视表定义部分。|
|数据透视表缓存记录(Pivot Table Cache Records)|包含数据透视表的基础数据。包中的每个数据透视表都有零个或一个这样的部分。|
|查询表(Query Table)|包含描述源表如何连接到外部数据源的信息，并定义从源刷新表时使用的属性。每个表可能有一个这样的部分。|
|共享字符串表(Shared String Table)|包含工作簿中任何工作表中出现的每个唯一字符串的一次出现。包装有一个这样的部分。|
|共享工作簿修订日志(Shared Workbook Revision Log)|包含有关在工作簿的工作表中对单个单元格执行的编辑的信息。每个编辑会话应该有一个这样的部分|
|共享工作簿用户数据(Shared Workbook User Data)|包含共享工作簿的所有用户的列表。包装包含零个或一个这样的部分|
|单细胞表定义(Single Cell Table Definition)|包含有关如何将自定义XML文件中的非重复元素映射到工作表中的单元格的信息。每个工作表可能有一个这样的部分。|
|样式(Styles)|包含工作簿中所有单元格的所有特征，包括数字和文本格式，对齐方式，字体，颜色和边框。包中只包含一个这样的部分。|
|表定义(Table Definition)|电子表格可以包含具有明确标记的列，行和数据区域的区域。这些区域称为表，每个工作表可以有零个或多个表。对于工作表中的每个表，必须有一个包含表描述的表定义部分。（表的数据存储在相应的工作表部分中。）|
|易失性依赖(Volatile Dependencies)|单元格可以包含实时数据公式，这些公式返回随时间变化的值，并且需要连接到工作簿之外的程序。如果这些程序不可用，公式可以使用存储在易失性依赖项部分中的信息。包装只能有一个这样的部件|
|工作簿(Workbook)|包含所有工作表的数据和引用。必须有一个且只有一个工作簿部分。|
|工作表(Worksheet)|包含给定工作表的所有数据，公式和特征。包中的每个工作表都有一个这样的部分。|


