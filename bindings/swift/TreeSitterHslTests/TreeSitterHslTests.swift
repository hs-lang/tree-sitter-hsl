import XCTest
import SwiftTreeSitter
import TreeSitterHsl

final class TreeSitterHslTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_hsl())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading HyperSpace Lang grammar")
    }
}
