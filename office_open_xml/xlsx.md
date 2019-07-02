## 参考文档
http://officeopenxml.com/anatomyofOOXML-xlsx.php

## xlsx Package Structure

`SpreadsheetML`或`.xlsx`文件是包含许多“部件”（通常为UTF-8或UTF-16编码）或XML文件的zip文件（包）。该包还可能包含其他媒体文件，如图片。
该结构根据OOXML标准ECMA-376第2部分中概述的开放式包装惯例进行组织。


只需解压缩.xlsx文件，即可查看包含SpreadsheetML文件的文件结构和文件。

```
 - _rels
 - docProps
 - xl
 - [Content_Types].xml
```

部件的数量和类型将根据电子表格中的内容而有所不同，但始终会有[Content_Types] .xml，一个或多个关系部件，工作簿部件和至少一个工作表。电子表格的核心数据包含在工作表部分中，在xsxl内容概述中有更详细的讨论(http://officeopenxml.com/SScontentOverview.php)。
