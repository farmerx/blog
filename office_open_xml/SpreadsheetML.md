
##  SpreadsheetML特定术语表

* AutoFilter  - 一种SpreadsheetML文档状态，其中仅显示某些行，通过过滤器确定应用于列的标准。
* Cell(单元格) - 行和列交叉处的位置，其中包含数字或文本数据或公式存储。 单元格可以具有许多特征，例如数字或文本格式，对齐方式，字体，颜色和边界等。
* Cell reference（单元格引用） - 使用其列和行标题的组合来指定单个单元格，如A13中所示，H19和BX1200。 公式向下复制时，公式中的相对单元格引用会自动更改
一列或一行。 绝对单元格引用是固定的。 公式时绝对引用不会改变从一个单元格复制到另一个单元格。 混合单元格引用具有绝对列和相对行，或绝对行和相对列。
* chart  - 数据的图形表示，例如，在条形图，列，线，饼图中。
* column  - 工作表中的任何垂直单元格集。 每列都有一个字母标题。 列被命名从A-Z开始，然后是AA-AZ，BA-BZ等。
* connection（连接） - 外部数据的方式，即数据存储在工作簿之外（在数据库中）例如，在Web服务器上 - 可以导入到工作表中。
* formula  - 计算值的方法。 一些公式是预定义的; 其他的是用户定义的。
* function  - 预定义的公式，例如AVERAGE，MAX，MIN和SUM。 一个函数需要一个或多个运作的参数，产生结果。 [注意：在公式= SUM（B1：B4）中，有一个参数，
B1：B4，其为细胞B1-B4的范围，包括端值。 结束说明]
* MDX  - 一种传递给OLAP提供程序的多维表达式语言。 解释的方法这是由服务器端OLAP提供程序实现定义的。
* OLAP  - 一种使用多维数据模型的在线分析处理数据库。
* pivottable（数据透视表） - 一种用于管理和分析存储在其他地方的相关数据的表。
* row  - 工作表中的任何水平单元格集。 每行都有一个数字标题。 行编号顺序，从1开始。
* table  - 一组矩形形状的相关行和列，可以对其进行排序，过滤和总计组。 通过将自动筛选器应用于一个或多个列，可以隐藏表中的行。
* workbook（工作簿） - 工作表的集合。
* worksheet（工作表） - 按行和列组织的二维单元格网格。

## 包结构（package structure）


## Part Summary

### Calculation Chain Part（计算链部分）


|    |         |  
|----------|:-------------:|
| Content Type |  application/vnd.openxmlformats-officedocument.spreadsheetml.calcChain+xml |
| Root Namespace |  http://purl.oclc.org/ooxml/spreadsheetml/main  | 
| Source Relationship |http://purl.oclc.org/ooxml/officeDocument/relationships/calcChain |

此部件类型的实例包含对工作簿中所有工作表中所有单元格的有序引用集其值由任何公式计算。 排序允许计算相互关联的单元格公式加载工作表以供使用时的正确顺序。
包裹不得超过一个计算链部分。 如果存在，该部分应成为目标工作手册部分的隐含关系（第12.3.23节）。

示例：以下工作簿部件关系项包含与计算链部件的关系，它存储在ZIP项calcChain.xml中：
```
.../xl/_rels/workbook.xml.rels
<Relationships xmlns="…">
 <Relationship Id="rId7"
 Type="http://…/calcChain" Target="calcChain.xml"/>
</Relationships>
```
Cells D8, E8, and F8 each contain a value that is the result of calculations that shall be performed inthe order E8, D8, F8:
```
<calcChain xmlns="…">
 <c r="E8" i="1"/>
 <c r="D8"/>
 <c r="F8" s="1"/>
</calcChain>
```
end example


计算链部分应位于包含关系部分的包装内（表示从语法上讲，Relationship元素的TargetMode属性应为Internal）。计算链部分不得与ECMA-376定义的任何部分具有隐式或显式关系。


### Chartsheet Part

|    |         |  
|----------|:-------------:|
| Content Type | application/vnd.openxmlformats-officedocument.spreadsheetml.chartsheet+xml |
| Root Namespace |  http://purl.oclc.org/ooxml/spreadsheetml/main  | 
| Source Relationship |http://purl.oclc.org/ooxml/officeDocument/relationships/chartsheet |




